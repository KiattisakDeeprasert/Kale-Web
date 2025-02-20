package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"mime/multipart"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = generateSecretKey()

func generateSecretKey() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("Error generating secret key:", err)
		return "defaultSecretKey"
	}
	return base64.StdEncoding.EncodeToString(key)
}

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type UploadHistory struct {
	Filename string    `json:"filename" bson:"filename"`
	Size     int64     `json:"size" bson:"size"`
	Uploaded time.Time `json:"uploaded" bson:"uploaded"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var historyCollection *mongo.Collection
var userCollection *mongo.Collection

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb+srv://nawinwin88:Nawinwin46@cluster0.xdqtr.mongodb.net/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("ไม่สามารถเชื่อมต่อกับ MongoDB ได้:", err)
		return
	}

	historyCollection = client.Database("uploadDB").Collection("uploadHistory")
	userCollection = client.Database("uploadDB").Collection("users")

	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/upload", authMiddleware(uploadHandler))
	http.HandleFunc("/history", authMiddleware(historyHandler))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		fmt.Println("กำลังปิดการทำงานอย่างสมบูรณ์...")
		cancel()
		client.Disconnect(ctx)
		os.Exit(0)
	}()

	// เซิร์ฟเวอร์
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := tokenHeader[len("Bearer "):]
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

// ตรวจสอบไฟล์
func validateFile(file multipart.File, size int64) error {
	// จำกัดขนาดไฟล์ไม่เกิน 10MB
	if size > 10*1024*1024 {
		return fmt.Errorf("ไฟล์มีขนาดใหญ่เกินไป")
	}

	// อ่านข้อมูลบางส่วนจากไฟล์เพื่อใช้ในการตรวจสอบประเภท
	buffer := make([]byte, 512) // ขนาดที่แนะนำสำหรับตรวจสอบประเภทไฟล์
	_, err := file.Read(buffer)
	if err != nil && err.Error() != "EOF" {
		return fmt.Errorf("ไม่สามารถอ่านไฟล์ได้")
	}

	// ตรวจสอบประเภทไฟล์จากข้อมูลที่อ่านมา
	mimeType := http.DetectContentType(buffer)
	if mimeType != "image/png" && mimeType != "image/jpeg" {
		return fmt.Errorf("ประเภทไฟล์ไม่ถูกต้อง")
	}

	// รีเซ็ตการอ่านไฟล์ไปที่จุดเริ่มต้นก่อนการอัปโหลด
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("ไม่สามารถรีเซ็ตการอ่านไฟล์ได้")
	}

	return nil
}

// อัปโหลดไฟล์
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "ไม่สามารถอ่านไฟล์ได้", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// ตรวจสอบไฟล์ก่อนการอัปโหลด
	if err := validateFile(file, header.Size); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filename := filepath.Join("uploads", header.Filename)
	out, err := os.Create(filename)
	if err != nil {
		http.Error(w, "ไม่สามารถสร้างไฟล์ได้", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "ไม่สามารถบันทึกไฟล์ได้", http.StatusInternalServerError)
		return
	}

	record := UploadHistory{
		Filename: header.Filename,
		Size:     header.Size,
		Uploaded: time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = historyCollection.InsertOne(ctx, record)
	if err != nil {
		http.Error(w, "ไม่สามารถบันทึกประวัติการอัปโหลดได้", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "ไฟล์อัปโหลดสำเร็จ",
		"filename": header.Filename,
	})
}

// ดูประวัติการอัปโหลด
func historyHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := historyCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "ไม่สามารถดึงข้อมูลประวัติได้", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var histories []UploadHistory
	if err = cursor.All(ctx, &histories); err != nil {
		http.Error(w, "ไม่สามารถอ่านข้อมูลประวัติได้", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(histories)
}

// ลงทะเบียน
func registerHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := userCollection.CountDocuments(ctx, bson.M{"username": user.Username})
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if count > 0 {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)

	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, "Cannot register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// Login
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq User
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	err = userCollection.FindOne(ctx, bson.M{"username": loginReq.Username}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

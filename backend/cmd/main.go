package main

import (
	"backend/internal/auth/delivery/http"
	"backend/internal/auth/usecase"
	"backend/internal/user/repository"
	"backend/pkg/database"
	"log"
)

func main() {

	client, err := database.ConnectMongoDB("mongodb+srv://nawinwin88:Nawinwin46@cluster0.xdqtr.mongodb.net/")
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("auth_db")

	// Create user repository and usecase
	userRepo := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUsecase(userRepo)
	authHandler := http.NewAuthHandler(authUsecase)

	// Register routes
	http.HandleFunc("/register", authHandler.Register)
	http.HandleFunc("/login", authHandler.Login)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

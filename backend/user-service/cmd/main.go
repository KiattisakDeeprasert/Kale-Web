package main

import (
	"context"
	"fmt"
	"myapp/internal/handler"
	"myapp/internal/repository"
	"myapp/internal/usecase"
	"myapp/pkg/jwtutil"

	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://your-uri"))
	userRepo := repository.NewUserRepository(client.Database("appDB").Collection("users"))
	jwtService := jwtutil.NewJWTService("your-secret-key")
	userUsecase := usecase.NewUserUsecase(userRepo, jwtService)
	userHandler := handler.NewUserHandler(userUsecase)

	http.HandleFunc("/register", userHandler.Register)
	http.HandleFunc("/login", userHandler.Login)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

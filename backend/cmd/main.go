package main

import (
	"backend/internal/auth/delivery"
	"backend/internal/auth/repository"
	"backend/internal/auth/usecase"
	"backend/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	userRepo := repository.NewUserRepository()
	authUsecase := usecase.NewAuthUsecase(userRepo)
	authHandler := delivery.NewAuthHandler(authUsecase)

	r := gin.Default()
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	r.Run(":8080")
}

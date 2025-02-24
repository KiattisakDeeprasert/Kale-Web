package repository

import (
	"backend/internal/auth/domain"
	"backend/pkg/database"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// บันทึกผู้ใช้ใหม่
func (r *UserRepository) CreateUser(ctx context.Context, user domain.User) error {
	_, err := database.DB.Collection("users").InsertOne(ctx, user)
	return err
}

// ค้นหาผู้ใช้ตามอีเมล
func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := database.DB.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

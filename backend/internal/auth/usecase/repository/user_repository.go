package mongodb

import (
	"context"

	"backend/internal/auth/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
	db *mongo.Database
}

// NewUserRepository เป็นตัวสร้าง repository
func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{db: db}
}

// FindUserByEmail ค้นหาผู้ใช้โดยอีเมล
func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	collection := r.db.Collection("users")
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

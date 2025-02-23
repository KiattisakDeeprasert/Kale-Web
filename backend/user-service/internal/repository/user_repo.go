package repository

import (
	"context"
	"errors"
	"myapp/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Collection) domain.UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := r.db.CountDocuments(ctx, bson.M{"username": user.Username})
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already exists")
	}

	_, err = r.db.InsertOne(ctx, user)
	return err
}

func (r *userRepo) GetUserByUsername(username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user domain.User
	err := r.db.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

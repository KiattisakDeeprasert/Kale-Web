package repository

import (
    "errors"
    "user-service/internal/entity"
)

type UserRepositoryMemory struct {
    users map[string]*entity.User
}

func NewUserRepositoryMemory() *UserRepositoryMemory {
    return &UserRepositoryMemory{users: make(map[string]*entity.User)}
}

func (r *UserRepositoryMemory) Create(user *entity.User) error {
    if _, exists := r.users[user.Email]; exists {
        return errors.New("user already exists")
    }
    r.users[user.Email] = user
    return nil
}

func (r *UserRepositoryMemory) GetByEmail(email string) (*entity.User, error) {
    user, exists := r.users[email]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}


package usecase

import (
	"backend/internal/auth/domain"
	"backend/internal/user/repository"
	"backend/pkg/utils"
	"context"
	"errors"
)

type AuthUsecase struct {
	userRepo repository.UserRepository
}

func NewAuthUsecase(userRepo repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{userRepo: userRepo}
}

func (u *AuthUsecase) Register(ctx context.Context, email, password string) error {
	_, err := u.userRepo.FindUserByEmail(ctx, email)
	if err == nil {
		return errors.New("email already exists")
	}
	user := &domain.User{Email: email, Password: password}
	return u.userRepo.CreateUser(ctx, user)
}

func (u *AuthUsecase) Login(ctx context.Context, email, password string) (*domain.User, error) {
	user, err := u.userRepo.FindUserByEmail(ctx, email)
	if err != nil || user.Password != password {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

func (u *AuthUsecase) GoogleLogin(ctx context.Context, token string) (*domain.User, error) {
	// Logic to verify Google token and fetch user data
}

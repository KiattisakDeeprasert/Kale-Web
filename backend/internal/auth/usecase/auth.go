package usecase

import (
	"backend/internal/auth/domain"
	"backend/internal/auth/repository"
	"backend/pkg/utils"
	"context"
	"errors"
)

type AuthUsecase struct {
	userRepo *repository.UserRepository
}

func NewAuthUsecase(userRepo *repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{userRepo: userRepo}
}


func (uc *AuthUsecase) Register(ctx context.Context, user domain.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return uc.userRepo.CreateUser(ctx, user)
}

// เข้าสู่ระบบ
func (uc *AuthUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := uc.userRepo.FindUserByEmail(ctx, email)
	if err != nil || !utils.CheckPassword(user.Password, password) {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

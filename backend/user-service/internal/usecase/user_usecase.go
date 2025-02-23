package usecase

import (
	"errors"
	"myapp/internal/domain"
	"myapp/pkg/hashutil"
)

type UserUsecase struct {
	userRepo   domain.UserRepository
	jwtService domain.JWTService
}

func NewUserUsecase(repo domain.UserRepository, jwtService domain.JWTService) *UserUsecase {
	return &UserUsecase{userRepo: repo, jwtService: jwtService}
}

func (u *UserUsecase) Register(username, password string) error {
	hashedPassword, err := hashutil.HashPassword(password)
	if err != nil {
		return err
	}

	user := &domain.User{
		Username: username,
		Password: hashedPassword,
	}
	return u.userRepo.CreateUser(user)
}

func (u *UserUsecase) Login(username, password string) (string, error) {
	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !hashutil.CheckPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	return u.jwtService.GenerateToken(user.Username)
}

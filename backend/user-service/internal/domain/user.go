package domain

import "time"


type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}


type UserRepository interface {
	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)
}

// JWT Service Interface
type JWTService interface {
	GenerateToken(username string) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
}

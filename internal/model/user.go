package model

import "github.com/golang-jwt/jwt/v5"

type CreateUser struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            int64
}

type UserClaims struct {
	id       int64
	Username string
	Email    string
	Role     int64
	jwt.Claims
}

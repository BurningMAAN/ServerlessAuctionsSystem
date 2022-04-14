package models

import "github.com/dgrijalva/jwt-go"

type AuthorizationConfig struct {
	Token string
}

type Claims struct {
	Username string `json:"username"`
	Role     `json:"role"`
	jwt.StandardClaims
}

type UserConfig struct {
	Name  string
	Token string
}

type Role string

const (
	UserRole     Role = "standard_user"
	BugalterRole Role = "bugalter"
)

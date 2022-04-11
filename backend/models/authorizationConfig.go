package models

import "github.com/dgrijalva/jwt-go"

type AuthorizationConfig struct {
	Token string
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type UserConfig struct {
	Name  string
	Token string
}

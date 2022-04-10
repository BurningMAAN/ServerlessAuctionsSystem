package authorizer

import (
	"auctionsPlatform/models"
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const (
	defaultExpirationTime int = 180
)

type userRepository interface {
	GetUserByUserName(ctx context.Context, userName string) (models.User, error)
}

type service struct {
	userRepository userRepository
}

func New(userRepository userRepository) *service {
	return &service{
		userRepository: userRepository,
	}
}

func (s *service) Authorize(ctx context.Context, userName, password string) (models.AuthorizationConfig, error) {
	user, err := s.userRepository.GetUserByUserName(ctx, userName)
	if err != nil {
		return models.AuthorizationConfig{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return models.AuthorizationConfig{}, fmt.Errorf("comparing error: %w", err)
	}

	expirationTime := time.Now().Add(time.Duration(defaultExpirationTime) * time.Minute)
	claims := &models.Claims{
		Username: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(user.Password)
	if err != nil {
		return models.AuthorizationConfig{}, fmt.Errorf("Signed string: %w", err)
	}

	return models.AuthorizationConfig{
		Token: tokenString,
	}, nil
}

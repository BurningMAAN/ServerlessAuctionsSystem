package user

import (
	"auctionsPlatform/models"
	"context"
)

type userRepository interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	GetUserByID(ctx context.Context, userID string) (models.User, error)
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

func (s *service) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	return s.userRepository.CreateUser(ctx, user)
}

func (s *service) GetUserByUserName(ctx context.Context, userName string) (models.User, error) {
	return s.userRepository.GetUserByUserName(ctx, userName)
}

func (s *service) GetUserByID(ctx context.Context, userID string) (models.User, error) {
	return s.userRepository.GetUserByID(ctx, userID)
}

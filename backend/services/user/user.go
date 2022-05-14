package user

import (
	"auctionsPlatform/models"
	"context"
	"encoding/json"
	"log"
)

type userRepository interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	GetUserByID(ctx context.Context, userID string) (models.User, error)
	GetUserByUserName(ctx context.Context, userName string) (models.User, error)
	UpdateUser(ctx context.Context, updateModel models.UserUpdate) error
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

func (s *service) UpdateUser(ctx context.Context, updateModel models.UserUpdate) error {
	updateBytes, _ := json.Marshal(updateModel)
	log.Print(string(updateBytes))
	user, err := s.userRepository.GetUserByUserName(ctx, updateModel.UserName)
	if err != nil {
		return err
	}

	if updateModel.Credit != nil {
		var updateCredit = *updateModel.Credit
		updateCredit += user.Credit
		updateModel.Credit = &updateCredit
	}

	return s.userRepository.UpdateUser(ctx, updateModel)
}

package user

import (
	"auctionsPlatform/models"
	"context"
)

type DB interface{}

type repository struct {
	DB DB
}

func New(db DB) *repository {
	return &repository{
		DB: db,
	}
}

type UserDB struct {
	PK       string // Example: User#{UserID}
	SK       string // Example: Metadata
	Password string
	GSI1PK   string // Example: User#{UserName}
	GSI1SK   string // Example: Metadata
}

func (r *repository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	return models.User{}, nil
}

func (r *repository) GetUserByID(ctx context.Context, userID string) (models.User, error) {
	return models.User{}, nil
}

func (r *repository) GetByUserName(ctx context.Context, userName string) (models.User, error) {
	return models.User{}, nil
}

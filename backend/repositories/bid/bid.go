package bid

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

type BidDB struct {
}

func (r *repository) CreateBid(ctx context.Context, auctionID string, bid models.Bid) (models.Bid, error) {
	return models.Bid{}, nil
}

func (r *repository) GetBidByID(ctx context.Context, auctionID, bidID string) (models.Bid, error) {
	return models.Bid{}, nil
}

func (r *repository) GetByIDAndAuctionID(ctx context.Context, auctionID, bidID string) (models.Bid, error) {
	return models.Bid{}, nil
}

func (r *repository) GetLatestAuctionBid(ctx context.Context, auctionID string) (*models.Bid, error) {
	return &models.Bid{}, nil
}

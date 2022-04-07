package auction

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

type AuctionDB struct {
	PK string // Example: Auction#{AuctionID}
	SK string // Example: Metadata
}

func (r *repository) CreateAuction(ctx context.Context, auction models.Auction) (models.Auction, error) {
	return models.Auction{}, nil
}

func (r *repository) GetAuctionByID(ctx context.Context, auctionID string) (models.Auction, error) {
	return models.Auction{}, nil
}

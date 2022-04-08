package auction

import (
	"auctionsPlatform/models"
	"context"
)

type auctionRepository interface {
	CreateAuction(ctx context.Context, auction models.Auction) (models.Auction, error)
	GetAuctionByID(ctx context.Context, auctionID string) (models.Auction, error)
}

type service struct {
	auctionRepository auctionRepository
}

func New(auctionRepository auctionRepository) *service {
	return &service{
		auctionRepository: auctionRepository,
	}
}

func (s *service) CreateAuction(ctx context.Context, auction models.Auction, itemID string) (models.Auction, error) {
	return s.auctionRepository.CreateAuction(ctx, auction)
}

func (s *service) GetAuctionByID(ctx context.Context, auctionID string) (models.Auction, error) {
	return s.auctionRepository.GetAuctionByID(ctx, auctionID)
}

package bid

import (
	"auctionsPlatform/errors"
	"auctionsPlatform/models"
	"context"
)

type bidRepository interface {
	CreateBid(ctx context.Context, auctionID string, bid models.Bid) (models.Bid, error)
	GetBidByID(ctx context.Context, bidID string) (models.Bid, error)
	GetLatestAuctionBid(ctx context.Context, auctionID string) (*models.Bid, error)
}

type auctionRepository interface {
	GetAuctionByID(ctx context.Context, auctionID string) (models.Auction, error)
}

type service struct {
	auctionRepository auctionRepository
	bidRepository     bidRepository
}

func New(auctionRepository auctionRepository, bidRepository bidRepository) *service {
	return &service{
		auctionRepository: auctionRepository,
		bidRepository:     bidRepository,
	}
}

func (s *service) PlaceBid(ctx context.Context, auctionID string, bid models.Bid) (models.Bid, error) {
	auction, err := s.auctionRepository.GetAuctionByID(ctx, auctionID)
	if err != nil {
		return models.Bid{}, err
	}

	if auction.IsFinished {
		return models.Bid{}, errors.ErrAuctionAlreadyFinished
	}

	switch auction.Type {
	case models.AuctionTypeAbsolute:
		err := s.handleAbsoluteBid(ctx, auction.ID, bid)
		if err != nil {
			return models.Bid{}, err
		}
	}

	return s.bidRepository.CreateBid(ctx, auction.ID, bid)
}

func (s *service) GetBidByID(ctx context.Context, bidID string) (models.Bid, error) {
	return s.bidRepository.GetBidByID(ctx, bidID)
}

package auction

import (
	"auctionsPlatform/errors"
	"auctionsPlatform/models"
	"auctionsPlatform/repositories/auction"
	"context"
)

type auctionRepository interface {
	CreateAuction(ctx context.Context, auction models.Auction) (models.Auction, error)
	GetAuctionByID(ctx context.Context, auctionID string) (models.Auction, error)
	GetAllAuctions(ctx context.Context, optFns ...func(*auction.OptionalGetParameters)) ([]models.Auction, error)
	FinishAuction(ctx context.Context, auctionID string) error
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

func (s *service) GetAuctions(ctx context.Context) ([]models.Auction, error) {
	return s.auctionRepository.GetAllAuctions(ctx)
}

func (s *service) FinishAuction(ctx context.Context, auctionID string) error {
	auction, err := s.auctionRepository.GetAuctionByID(ctx, auctionID)
	if err != nil {
		return err
	}

	if auction.IsFinished {
		return errors.ErrAuctionAlreadyFinished
	}

	return s.auctionRepository.FinishAuction(ctx, auctionID)
}

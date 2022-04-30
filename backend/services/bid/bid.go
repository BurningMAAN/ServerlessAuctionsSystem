package bid

import (
	"auctionsPlatform/errors"
	"auctionsPlatform/models"
	bidRepo "auctionsPlatform/repositories/bid"
	"context"
	"fmt"
	"time"
)

type bidRepository interface {
	CreateBid(ctx context.Context, auctionID string, bid models.Bid) (models.Bid, error)
	GetBidByID(ctx context.Context, bidID string) (models.Bid, error)
	GetLatestAuctionBids(ctx context.Context, auctionID string, optFns ...func(*bidRepo.OptionalGetParameters)) ([]models.Bid, error)
}

type auctionRepository interface {
	GetAuctionByID(ctx context.Context, auctionID string) (models.Auction, error)
	UpdateAuctionWorker(ctx context.Context, auctionID string, updateModel models.AuctionWorkerUpdateModel) error
}

type userRepository interface {
	GetUserByID(ctx context.Context, userID string) (models.User, error)
}

type service struct {
	auctionRepository auctionRepository
	bidRepository     bidRepository
	userRepository    userRepository
}

func New(auctionRepository auctionRepository, bidRepository bidRepository, userRepository userRepository) *service {
	return &service{
		auctionRepository: auctionRepository,
		bidRepository:     bidRepository,
		userRepository:    userRepository,
	}
}

func (s *service) PlaceBid(ctx context.Context, auctionID string, bid models.Bid) (models.Bid, error) {
	auction, err := s.auctionRepository.GetAuctionByID(ctx, auctionID)
	if err != nil {
		return models.Bid{}, fmt.Errorf("auction err: %s", err.Error())
	}

	user, err := s.userRepository.GetUserByID(ctx, bid.UserID)
	if err != nil {
		return models.Bid{}, fmt.Errorf("user err: %s", err.Error())
	}

	if auction.CreatorID == user.ID {
		return models.Bid{}, errors.ErrAuctionUserBidUserMatch
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

	placedBid, err := s.bidRepository.CreateBid(ctx, auction.ID, bid)
	if err != nil {
		return models.Bid{}, err
	}

	newEndDate := time.Now().Add(33 * time.Second)
	err = s.auctionRepository.UpdateAuctionWorker(ctx, auction.ID, models.AuctionWorkerUpdateModel{
		EndDate: &newEndDate,
	})

	if err != nil {
		return placedBid, err
	}

	return placedBid, err
}

func (s *service) GetBidByID(ctx context.Context, bidID string) (models.Bid, error) {
	return s.bidRepository.GetBidByID(ctx, bidID)
}

func (s *service) GetLatestAuctionBids(ctx context.Context, auctionID string, optFns ...func(*bidRepo.OptionalGetParameters)) ([]models.Bid, error) {
	auction, err := s.auctionRepository.GetAuctionByID(ctx, auctionID)
	if err != nil {
		return []models.Bid{}, err
	}

	return s.bidRepository.GetLatestAuctionBids(ctx, auction.ID)
}

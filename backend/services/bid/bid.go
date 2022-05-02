package bid

import (
	"auctionsPlatform/errors"
	"auctionsPlatform/models"
	bidRepo "auctionsPlatform/repositories/bid"
	"context"
	"fmt"
	"log"
	"time"
)

type bidRepository interface {
	CreateBid(ctx context.Context, auctionID string, bid models.Bid) (models.Bid, error)
	GetBidByID(ctx context.Context, bidID string) (models.Bid, error)
	GetLatestAuctionBids(ctx context.Context, auctionID string, optFns ...func(*bidRepo.OptionalGetParameters)) ([]models.Bid, error)
}

type auctionRepository interface {
	GetAuctionByID(ctx context.Context, auctionID string) (models.Auction, error)
	UpdateAuctionEndDate(ctx context.Context, auctionID string, endDate time.Time) error
}

type userRepository interface {
	GetUserByID(ctx context.Context, userID string) (models.User, error)
}

type eventRepository interface {
	UpdateEventRule(ctx context.Context, auctionID string) error
	CreateBidEvent(ctx context.Context, auctionID string) error
}

type service struct {
	auctionRepository auctionRepository
	bidRepository     bidRepository
	userRepository    userRepository
	eventRepository   eventRepository
}

func New(auctionRepository auctionRepository, bidRepository bidRepository, userRepository userRepository, eventRepository eventRepository) *service {
	return &service{
		auctionRepository: auctionRepository,
		bidRepository:     bidRepository,
		userRepository:    userRepository,
		eventRepository:   eventRepository,
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

	// sitas turi but workeryje
	// err = s.auctionRepository.UpdateAuctionEndDate(ctx, auctionID, time.Now().Add(33*time.Second))
	// if err != nil {
	// 	return models.Bid{}, err
	// }

	// err = s.eventRepository.UpdateEventRule(ctx, auction.ID)
	// if err != nil {
	// 	log.Printf("failed to place bid on auctionID: %s, error: %s", auction.ID, err.Error())
	// 	return placedBid, nil
	// }

	err = s.eventRepository.CreateBidEvent(ctx, auction.ID)
	if err != nil {
		log.Printf("failed to create bid entry for auction ID: %s, error: %s", auction.ID, err.Error())
		return placedBid, nil
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

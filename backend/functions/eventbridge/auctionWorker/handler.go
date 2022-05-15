package main

import (
	"auctionsPlatform/models"
	bidRepo "auctionsPlatform/repositories/bid"
	"context"
	"encoding/json"
	"log"
	"time"
)

type auctionRepository interface {
	UpdateAuctionStage(ctx context.Context, auctionID string, stage string) error
	UpdateAuctionEndDate(ctx context.Context, auctionID string, endDate time.Time) error
	UpdateAuction(ctx context.Context, auctionID string, update models.AuctionUpdate) error
}

type eventRepository interface {
	UpdateEventRule(ctx context.Context, auctionID string, newDate time.Time) error
	DeleteEventRule(ctx context.Context, auctionID string) error
}

type userRepository interface {
	UpdateUser(ctx context.Context, updateModel models.UserUpdate) error
	GetUserByUserName(ctx context.Context, userName string) (models.User, error)
}

type bidRepository interface {
	GetLatestAuctionBids(ctx context.Context, auctionID string, optFns ...func(*bidRepo.OptionalGetParameters)) ([]models.Bid, error)
}

type handler struct {
	auctionRepo     auctionRepository
	eventRepository eventRepository
	userRepository  userRepository
	bidRepository   bidRepository
}

func (h *handler) HandleAuction(ctx context.Context, event models.AuctionEvent) error {
	eventBytes, _ := json.Marshal(event)
	log.Print(string(eventBytes))
	switch event.Stage {
	case "STAGE_ACCEPTING_BIDS":
		err := h.auctionRepo.UpdateAuctionStage(ctx, event.AuctionID, "STAGE_AUCTION_ONGOING")
		if err != nil {
			log.Print(err.Error())
			return err
		}
		newEndTime := time.Now().Add(60 * time.Second)
		err = h.auctionRepo.UpdateAuctionEndDate(ctx, event.AuctionID, newEndTime)
		if err != nil {
			return err
		}

		err = h.eventRepository.UpdateEventRule(ctx, event.AuctionID, newEndTime)
		if err != nil {
			log.Print(err.Error())
			return err
		}
	case "STAGE_AUCTION_ONGOING":
		err := h.auctionRepo.UpdateAuctionStage(ctx, event.AuctionID, "STAGE_AUCTION_FINISHED")
		if err != nil {
			log.Print(err.Error())
			return err
		}

		err = h.eventRepository.DeleteEventRule(ctx, event.AuctionID)
		if err != nil {
			return err
		}

		bids, err := h.bidRepository.GetLatestAuctionBids(ctx, event.AuctionID)
		if err != nil {
			return err
		}

		if len(bids) > 0 {
			user, err := h.userRepository.GetUserByUserName(ctx, bids[0].UserID)
			if err != nil {
				return err
			}

			newCreditBalance := user.Credit - bids[0].Value
			if newCreditBalance >= 0 {
				err = h.auctionRepo.UpdateAuction(ctx, event.AuctionID, models.AuctionUpdate{
					WinnerID: &bids[0].UserID,
				})
				if err != nil {
					return err
				}

				err = h.userRepository.UpdateUser(ctx, models.UserUpdate{
					UserName: bids[0].UserID,
					Credit:   &newCreditBalance,
				})
				log.Printf("new credit balance: %v", newCreditBalance)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

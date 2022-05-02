package main

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"
	"context"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

type auctionRepository interface {
	UpdateAuctionStage(ctx context.Context, auctionID string, stage string) error
}

type eventRepository interface {
	UpdateEventRule(ctx context.Context, auctionID string) error
	DeleteEventRule(ctx context.Context, auctionID string) error
}

type handler struct {
	auctionRepo     auctionRepository
	eventRepository eventRepository
}

func (h *handler) HandleAuction(ctx context.Context, event interface{}) error {
	switch v := event.(type) {
	case models.AuctionEvent:
		switch v.Stage {
		case "STAGE_ACCEPTING_BIDS":
			err := h.auctionRepo.UpdateAuctionStage(ctx, v.AuctionID, "STAGE_AUCTION_ONGOING")
			if err != nil {
				log.Print(err.Error())
				return err
			}

			err = h.eventRepository.UpdateEventRule(ctx, v.AuctionID)
			if err != nil {
				log.Print(err.Error())
				return err
			}
		case "STAGE_AUCTION_ONGOING":
			err := h.auctionRepo.UpdateAuctionStage(ctx, v.AuctionID, "STAGE_AUCTION_FINISHED")
			if err != nil {
				log.Print(err.Error())
				return err
			}

			err = h.eventRepository.DeleteEventRule(ctx, v.AuctionID)
			if err != nil {
				return err
			}
		}
	case events.DynamoDBStreamRecord:
		bidID := utils.Extract("Bid", v.Keys["PK"].String())
		if len(bidID) <= 0 {
			return nil
		}

		auctionID := utils.Extract("Auction", v.NewImage["GSI1PK"].String())
		if len(auctionID) <= 0 {
			return errors.New("failed to retrieve auctionID for bid")
		}

		log.Print(bidID, auctionID)
	}
	return nil
}

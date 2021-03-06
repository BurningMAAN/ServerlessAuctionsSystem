package main

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"
	"context"
	"errors"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type auctionRepository interface {
	UpdateAuctionStage(ctx context.Context, auctionID string, stage string) error
	GetAuctionByID(ctx context.Context, auctionID string) (models.Auction, error)
	UpdateAuctionEndDate(ctx context.Context, auctionID string, endDate time.Time) error
}

type eventRepository interface {
	UpdateEventRule(ctx context.Context, auctionID string, newDate time.Time) error
	DeleteEventRule(ctx context.Context, auctionID string) error
}

type handler struct {
	auctionRepo     auctionRepository
	eventRepository eventRepository
}

func (h *handler) HandleBid(ctx context.Context, event events.DynamoDBEvent) error {
	bidID := utils.Extract("Bid", event.Records[0].Change.Keys["PK"].String())
	if len(bidID) == 0 {
		return nil
	}

	auctionID := utils.Extract("Auction", event.Records[0].Change.NewImage["GSI1PK"].String())
	if len(auctionID) == 0 {
		log.Print("failed to parse auctionID")
		return errors.New("failed to parse auctionID")
	}

	auction, err := h.auctionRepo.GetAuctionByID(ctx, auctionID)
	if err != nil {
		return err
	}

	if auction.Stage == "STAGE_AUCTION_ONGOING" {
		newEndTime := time.Now().Add(time.Minute)
		log.Printf("new end time: %s", newEndTime.String())
		err := h.auctionRepo.UpdateAuctionEndDate(ctx, auction.ID, newEndTime)
		if err != nil {
			return err
		}

		err = h.eventRepository.UpdateEventRule(ctx, auction.ID, newEndTime)
		if err != nil {
			return err
		}

	}

	return nil
}

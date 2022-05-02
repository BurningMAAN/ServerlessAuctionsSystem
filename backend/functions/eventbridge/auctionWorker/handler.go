package main

import (
	"auctionsPlatform/models"
	"context"
	"log"
	"time"
)

type auctionRepository interface {
	UpdateAuctionStage(ctx context.Context, auctionID string, stage string) error
}

type eventRepository interface {
	UpdateEventRule(ctx context.Context, auctionID string, newDate time.Time) error
	DeleteEventRule(ctx context.Context, auctionID string) error
}

type handler struct {
	auctionRepo     auctionRepository
	eventRepository eventRepository
}

func (h *handler) HandleAuction(ctx context.Context, event models.AuctionEvent) error {
	switch event.Stage {
	case "STAGE_ACCEPTING_BIDS":
		err := h.auctionRepo.UpdateAuctionStage(ctx, event.AuctionID, "STAGE_AUCTION_ONGOING")
		if err != nil {
			log.Print(err.Error())
			return err
		}

		newEndDate := time.Now().Add(33 * time.Second)
		err = h.eventRepository.UpdateEventRule(ctx, event.AuctionID, newEndDate)
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
	}
	return nil
}

package main

import (
	"context"
	"log"
	"time"
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

type AuctionEvent struct {
	AuctionID string    `json:"id"`
	Stage     string    `json:"stage"`
	EndDate   time.Time `json:"endDate"`
}

func (h *handler) HandleAuction(ctx context.Context, event AuctionEvent) error {
	log.Printf("got event: %s", event)
	switch event.Stage {
	case "STAGE_ACCEPTING_BIDS":
		err := h.auctionRepo.UpdateAuctionStage(ctx, event.AuctionID, "STAGE_AUCTION_ONGOING")
		if err != nil {
			log.Print(err.Error())
			return err
		}

		err = h.eventRepository.UpdateEventRule(ctx, event.AuctionID)
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

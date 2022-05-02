package main

import (
	"context"
	"encoding/json"
	"log"
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

func (h *handler) HandleAuction(ctx context.Context, gotEvent interface{}) error {
	eventBytes, _ := json.Marshal(gotEvent)
	log.Printf("got event: %s", string(eventBytes))

	// event := models.AuctionEvent{}
	// err := json.Unmarshal([]byte(gotEvent), )
	// switch event.Stage {
	// case "STAGE_ACCEPTING_BIDS":
	// 	err := h.auctionRepo.UpdateAuctionStage(ctx, event.AuctionID, "STAGE_AUCTION_ONGOING")
	// 	if err != nil {
	// 		log.Print(err.Error())
	// 		return err
	// 	}

	// 	err = h.eventRepository.UpdateEventRule(ctx, event.AuctionID)
	// 	if err != nil {
	// 		log.Print(err.Error())
	// 		return err
	// 	}
	// case "STAGE_AUCTION_ONGOING":
	// 	err := h.auctionRepo.UpdateAuctionStage(ctx, event.AuctionID, "STAGE_AUCTION_FINISHED")
	// 	if err != nil {
	// 		log.Print(err.Error())
	// 		return err
	// 	}

	// 	err = h.eventRepository.DeleteEventRule(ctx, event.AuctionID)
	// 	if err != nil {
	// 		return err
	// 	}
	// case "AUCTION_BID_PLACED":
	// 	// paupdeitinam endDate auctionDB ir atnaujinam data auction controller event'o
	// }
	return nil
}

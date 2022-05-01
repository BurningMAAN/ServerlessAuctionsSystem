package main

import (
	"context"
	"time"
)

type auctionRepository interface {
	UpdateAuctionStage(ctx context.Context, auctionID string, stage string) error
}

type eventRepository interface {
	UpdateEventRule(ctx context.Context, auctionID string) error
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
	switch event.Stage {
	case "STAGE_ACCEPTING_BIDS":
		err := h.auctionRepo.UpdateAuctionStage(ctx, event.AuctionID, "STAGE_AUCTION_ONGOING")
		if err != nil {
			return err
		}

		err = h.eventRepository.UpdateEventRule(ctx, event.AuctionID)
		if err != nil {
			return err
		}

	}
	// for _, eventRecord := range event.Records {
	// 	if eventRecord.EventName == "REMOVE" {
	// 		record := unmarshalEvent(eventRecord)

	// 		switch record.Status {
	// 		case "STATUS_ACCEPTING_BIDS":
	// 			// Create Worker entity with startDate = null, endDate = currentTime + 33s
	// 			endDate := time.Now().Add(33 * time.Second)
	// 			err := h.auctionRepo.CreateAuctionWorker(ctx, record.AuctionID, "STATUS_AUCTION_ONGOING", endDate)
	// 			if err != nil {
	// 				log.Print(err.Error())
	// 			}

	// 		case "STATUS_AUCTION_ONGOING":
	// 			err := h.auctionRepo.FinishAuction(ctx, record.AuctionID)
	// 			if err != nil {
	// 				log.Print(err.Error())
	// 			}
	// 		default:
	// 			log.Printf("unsupported entity status: %s", record.Status)
	// 		}
	// 	}
	// }
	return nil
}

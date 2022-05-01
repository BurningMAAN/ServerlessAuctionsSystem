package main

import (
	"auctionsPlatform/utils"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type auctionRepository interface {
	CreateAuctionWorker(ctx context.Context, auctionID string, status string, endDate time.Time) error
	FinishAuction(ctx context.Context, auctionID string) error
}

type DynamoDBEvent struct {
	EventName string `json:"eventName"`
}

type Record struct {
	AuctionID      string    `json:"auctionId"`
	Status         string    `json:"status"`
	AuctionEndDate time.Time `json:"auctionEndDate"`
}

type handler struct {
	auctionRepo auctionRepository
}

type AuctionEvent struct {
	AuctionID string    `json:"id"`
	Stage     string    `json:"stage"`
	EndDate   time.Time `json:"endDate"`
}

func (h *handler) HandleAuction(ctx context.Context, event AuctionEvent) {
	eventJSON, _ := json.Marshal(event)
	log.Print(string(eventJSON))
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
}

func unmarshalEvent(eventRecord events.DynamoDBEventRecord) Record {
	pk := eventRecord.Change.Keys["PK"].String()
	status := eventRecord.Change.OldImage["Status"].String()
	auctionStartDate, err := time.Parse(time.RFC3339, eventRecord.Change.OldImage["StartDate"].String())
	if err != nil {
		log.Printf("Nepavyko patraukt datos, gavom data: %s, err: %s", auctionStartDate.String(), err.Error())
	}

	auctionEndDate, err := time.Parse(time.RFC3339, eventRecord.Change.OldImage["EndDate"].String())
	if err != nil {
		log.Printf("Nepavyko patraukt datos, gavom data: %s, err: %s", auctionEndDate.String(), err.Error())
	}

	return Record{
		AuctionID:      utils.Extract("AuctionWorker", pk),
		Status:         status,
		AuctionEndDate: auctionEndDate,
	}
}

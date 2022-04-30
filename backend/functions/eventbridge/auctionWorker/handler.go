package main

import (
	"auctionsPlatform/utils"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type DynamoDBEvent struct {
	EventName string `json:"eventName"`
}

type Record struct {
	AuctionID      string    `json:"auctionId"`
	Status         string    `json:"status"`
	AuctionEndDate time.Time `json:"auctionEndDate"`
}

func HandleAuction(ctx context.Context, event events.DynamoDBEvent) {
	for _, eventRecord := range event.Records {
		pk := eventRecord.Change.Keys["PK"].String()
		status := eventRecord.Change.OldImage["Status"].String()
		auctionEndDate, err := time.Parse(time.RFC3339, eventRecord.Change.OldImage["EndDate"].String())
		if err != nil {
			panic("invalidi data")
		}

		transformedRecord := Record{
			AuctionID:      utils.Extract("Auction", pk),
			Status:         status,
			AuctionEndDate: auctionEndDate,
		}

		jsonas, err := json.Marshal(transformedRecord)
		if err != nil {
			panic(err)
		}

		log.Print(jsonas)
	}
}

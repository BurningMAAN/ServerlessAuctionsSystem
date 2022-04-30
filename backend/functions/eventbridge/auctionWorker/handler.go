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
	AuctionID        string    `json:"auctionId"`
	Status           string    `json:"status"`
	AuctionStartDate time.Time `json:"auctionStartDate"`
	AuctionEndDate   time.Time `json:"auctionEndDate"`
}

func HandleAuction(ctx context.Context, event events.DynamoDBEvent) {
	eventJSON, _ := json.Marshal(event)
	log.Print(string(eventJSON))
	for _, eventRecord := range event.Records {
		pk := eventRecord.Change.Keys["PK"].String()
		status := eventRecord.Change.OldImage["Status"].String()
		auctionEndDate, err := time.Parse(time.RFC3339, eventRecord.Change.OldImage["EndDate"].String())
		if err != nil {
			log.Printf("Nepavyko patraukt datos, gavom data: %s, err: %s", auctionEndDate.String(), err.Error())
		}

		transformedRecord := Record{
			AuctionID:      utils.Extract("AuctionWorker", pk),
			Status:         status,
			AuctionEndDate: auctionEndDate,
		}

		jsonas, err := json.Marshal(transformedRecord)
		if err != nil {
			panic(err)
		}

		log.Print(string(jsonas))
	}
}

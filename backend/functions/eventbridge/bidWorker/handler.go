package main

import (
	"context"
	"encoding/json"
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

func (h *handler) HandleBid(ctx context.Context, event events.DynamoDBEvent) error {
	eventBytes, _ := json.Marshal(event)
	log.Print(string(eventBytes))
	return nil
}

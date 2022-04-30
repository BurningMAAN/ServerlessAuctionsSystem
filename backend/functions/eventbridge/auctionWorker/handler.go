package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func HandleAuction(ctx context.Context, event events.DynamoDBStreamRecord) {
	log.Print(event)
}

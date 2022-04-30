package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func HandleAuction(ctx context.Context, event events.DynamoDBStreamRecord) {
	testString, _ := json.Marshal(event)
	log.Print(string(testString))
}

package main

import (
	"auctionsPlatform/models"
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type auctionService interface {
	CreateAuction(ctx context.Context, auction models.Auction, itemID string) (models.Auction, error)
}

type handler struct {
	auctionService auctionService
}

func (h *handler) CreateAuction(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "WHADDUP I AM WORKING",
	}, nil
}

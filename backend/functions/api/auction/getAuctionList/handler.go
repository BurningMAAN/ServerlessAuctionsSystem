package main

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type request struct {
	Limit string `json:"limit"`
}

type response struct {
	Auctions []auction `json:"auctions"`
}

type auction struct {
	ID           string    `json:"id"`
	AuctionDate  time.Time `json:"auctionDate"`
	BuyoutPrice  *float64  `json:"buyoutPrice"`
	AuctionType  string    `json:"auctionType"`
	BidIncrement float64   `json:"BidIncrement"`
	CreatorID    string    `json:"creatorId"`
}

type auctionService interface {
	GetAuctions(ctx context.Context) ([]models.Auction, error)
}

type handler struct {
	auctionService auctionService
}

func (h *handler) GetAuction(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	auctions, err := h.auctionService.GetAuctions(ctx)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	resp := auctionsToResponse(auctions)
	respBody, err := json.Marshal(resp)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		StatusCode: http.StatusOK,
		Body:       string(respBody),
	}, nil
}

func auctionsToResponse(auctions []models.Auction) response {
	auctionsList := []auction{}
	for _, auctionItem := range auctions {
		auctionsList = append(auctionsList, auction{
			ID:           auctionItem.ID,
			AuctionDate:  auctionItem.StartDate,
			AuctionType:  string(auctionItem.Type),
			BuyoutPrice:  auctionItem.BuyoutPrice,
			BidIncrement: auctionItem.BidIncrement,
			CreatorID:    auctionItem.CreatorID,
		})
	}

	return response{auctionsList}
}

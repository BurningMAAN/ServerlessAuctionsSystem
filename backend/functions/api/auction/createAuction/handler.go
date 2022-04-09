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
	ItemID       string    `json:"itemId"`
	AuctionDate  time.Time `json:"auctionDate"`
	BuyoutPrice  *float64  `json:"buyoutPrice"`
	AuctionType  string    `json:"auctionType"`
	BidIncrement float64   `json:"BidIncrement"`
	CreatorID    string    `json:"creatorId"`
}

type response struct {
	ID           string    `json:"id"`
	AuctionDate  time.Time `json:"auctionDate"`
	BuyoutPrice  *float64  `json:"buyoutPrice"`
	AuctionType  string    `json:"auctionType"`
	BidIncrement float64   `json:"BidIncrement"`
	CreatorID    string    `json:"creatorId"`
}

type auctionService interface {
	CreateAuction(ctx context.Context, auction models.Auction, itemID string) (models.Auction, error)
}

type handler struct {
	auctionService auctionService
}

func (h *handler) CreateAuction(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req := request{}
	err := json.Unmarshal([]byte(event.Body), &req)
	if err != nil {
		return utils.InternalError()
	}

	auction, err := h.auctionService.CreateAuction(ctx, models.Auction{
		Type:         models.AuctionType(req.AuctionType),
		StartDate:    req.AuctionDate,
		BidIncrement: req.BidIncrement,
		CreatorID:    req.CreatorID,
	}, req.ItemID)
	if err != nil {
		return utils.InternalError()
	}

	respBody, err := json.Marshal(response{
		ID:           auction.ID,
		AuctionDate:  auction.StartDate,
		BuyoutPrice:  auction.BuyoutPrice,
		AuctionType:  string(auction.Type),
		BidIncrement: auction.BidIncrement,
		CreatorID:    auction.CreatorID,
	})
	if err != nil {
		return utils.InternalError()
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		StatusCode: http.StatusCreated,
		Body:       string(respBody),
	}, nil
}

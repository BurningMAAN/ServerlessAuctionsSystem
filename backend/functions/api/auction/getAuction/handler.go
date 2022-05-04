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
	AuctionID string
}

type response struct {
	ID           string    `json:"id"`
	AuctionDate  time.Time `json:"auctionDate"`
	BuyoutPrice  *float64  `json:"buyoutPrice"`
	AuctionType  string    `json:"auctionType"`
	BidIncrement float64   `json:"bidIncrement"`
	CreatorID    string    `json:"creatorId"`
	IsFinished   bool      `json:"isFinished"`
	ItemID       string    `json:"itemId"`
	Stage        string    `json:"stage"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
}

type auctionService interface {
	GetAuctionByID(ctx context.Context, auctionID string) (models.Auction, error)
}

type handler struct {
	auctionService auctionService
}

func (h *handler) GetAuction(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if len(event.PathParameters["auctionId"]) == 0 {
		return utils.InternalError("not provided auctionId")
	}

	req := request{
		AuctionID: event.PathParameters["auctionId"],
	}

	auction, err := h.auctionService.GetAuctionByID(ctx, req.AuctionID)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	respBody, err := json.Marshal(response{
		ID:           auction.ID,
		AuctionDate:  auction.StartDate,
		BuyoutPrice:  auction.BuyoutPrice,
		AuctionType:  string(auction.Type),
		BidIncrement: auction.BidIncrement,
		CreatorID:    auction.CreatorID,
		ItemID:       auction.ItemID,
		Stage:        auction.Stage,
		StartDate:    auction.StartDate,
		EndDate:      auction.EndDate,
	})
	if err != nil {
		return utils.InternalError(err.Error())
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: http.StatusOK,
		Body:       string(respBody),
	}, nil
}

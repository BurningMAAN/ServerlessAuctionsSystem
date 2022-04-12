package main

import (
	"auctionsPlatform/models"
	bidRepo "auctionsPlatform/repositories/bid"
	"auctionsPlatform/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type request struct {
	AuctionID string
}

type response struct {
	Bids []bidList `json:"bids"`
}

type bidList struct {
	ID        string    `json:"id"`
	AuctionID string    `json:"auctionId"`
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
	UserID    string    `json:"userId"`
}
type bidService interface {
	GetLatestAuctionBids(ctx context.Context, auctionID string, optFns ...func(*bidRepo.OptionalGetParameters)) ([]models.Bid, error)
}

type handler struct {
	bidService bidService
}

func (h *handler) PlaceBid(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	accessToken := event.Headers["access_token"]
	if len(accessToken) <= 0 {
		return utils.InternalError("token not provided")
	}

	_, err := utils.GetUserConfig(accessToken)
	if err != nil {
		return utils.InternalError(fmt.Sprintf("get user config err: %s", err.Error()))
	}

	if len(event.PathParameters["auctionId"]) <= 0 {
		return utils.InternalError("auctionID not provided")
	}
	req := request{
		AuctionID: event.PathParameters["auctionId"],
	}

	bids, err := h.bidService.GetLatestAuctionBids(ctx, req.AuctionID)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	response := response{}
	for _, bid := range bids {
		response.Bids = append(response.Bids, bidList{
			ID:        bid.ID,
			AuctionID: bid.AuctionID,
			Timestamp: bid.Timestamp,
			Value:     bid.Value,
		})
	}
	respBody, err := json.Marshal(response)
	if err != nil {
		return utils.InternalError(fmt.Sprintf("marshal err: %s", err.Error()))
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "*",
		},
		StatusCode: http.StatusOK,
		Body:       string(respBody),
	}, nil
}

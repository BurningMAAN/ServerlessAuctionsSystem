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
	Value     float64 `json:"value"`
}

type response struct {
	ID        string    `json:"id"`
	AuctionID string    `json:"auctionId"`
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
}

type bidService interface {
	PlaceBid(ctx context.Context, auctionID string, bid models.Bid) (models.Bid, error)
}

type handler struct {
	bidService bidService
}

func (h *handler) PlaceBid(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	accessToken := event.Headers["access_token"]
	if len(accessToken) <= 0 {
		return utils.InternalError("token not provided")
	}

	userConfig, err := utils.GetUserConfig(accessToken)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	if len(event.PathParameters["auctionId"]) <= 0 {
		return utils.InternalError("auctionID not provided")
	}

	req := request{
		AuctionID: event.PathParameters["auctionId"],
	}

	err = json.Unmarshal([]byte(event.Body), &req)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	bid, err := h.bidService.PlaceBid(ctx, req.AuctionID, models.Bid{
		Value:     req.Value,
		AuctionID: req.AuctionID,
		Timestamp: time.Now(),
		UserID:    userConfig.Name,
	})
	if err != nil {
		return utils.InternalError(err.Error())
	}

	respBody, err := json.Marshal(response{
		ID:        bid.ID,
		AuctionID: bid.AuctionID,
		Timestamp: bid.Timestamp,
		Value:     bid.Value,
	})
	if err != nil {
		return utils.InternalError(err.Error())
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "*",
		},
		StatusCode: http.StatusCreated,
		Body:       string(respBody),
	}, nil
}

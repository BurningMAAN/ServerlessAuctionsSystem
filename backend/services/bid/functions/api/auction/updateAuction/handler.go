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
	AuctionID    string
	BuyoutPrice  float64   `json:"buyoutPrice"`
	StartDate    time.Time `json:"startDate"`
	BidIncrement float64   `json:"bidIncrement"`
	Type         string    `json:"type"`
}

type auctionService interface {
	UpdateAuction(ctx context.Context, auctionID string, update models.AuctionUpdate) error
}

type handler struct {
	auctionService auctionService
}

func (h *handler) CreateAuction(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if len(event.Headers["access_token"]) == 0 {
		return utils.InternalError("no auth token provided")
	}

	if len(event.PathParameters["auctionId"]) == 0 {
		return utils.InternalError("not provided auctionId")
	}

	// userConfig, err := utils.GetUserConfig(event.Headers["access_token"])
	// if err != nil {
	// 	return utils.InternalError(err.Error())
	// }

	req := request{}
	req.AuctionID = event.PathParameters["auctionId"]
	err := json.Unmarshal([]byte(event.Body), &req)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	err = h.auctionService.UpdateAuction(ctx, req.AuctionID, models.AuctionUpdate{
		BuyoutPrice:  &req.BuyoutPrice,
		StartDate:    &req.StartDate,
		BidIncrement: &req.BidIncrement,
		Type:         &req.Type,
	})
	if err != nil {
		return utils.InternalError(err.Error())
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: http.StatusNoContent,
	}, nil
}

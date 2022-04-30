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
	ID      string    `json:"id"`
	Status  string    `json:"status"`
	EndDate time.Time `json:"endDate"`
}

type auctionService interface {
	GetAuctionWorker(ctx context.Context, auctionID string) (models.AuctionWorker, error)
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

	auctionWorker, err := h.auctionService.GetAuctionWorker(ctx, req.AuctionID)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	respBody, err := json.Marshal(response{
		ID:      auctionWorker.AuctionID,
		Status:  auctionWorker.Status,
		EndDate: auctionWorker.EndDate,
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

package main

import (
	"auctionsPlatform/utils"
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type request struct {
	AuctionID string
}

type auctionService interface {
	DeleteAuction(ctx context.Context, auctionID string) error
}

type handler struct {
	auctionService auctionService
}

func (h *handler) DeleteAuction(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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

	req := request{
		AuctionID: event.PathParameters["auctionId"],
	}
	err := h.auctionService.DeleteAuction(ctx, req.AuctionID)
	if err != nil {
		return utils.InternalError(err.Error())
	}
	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: http.StatusCreated,
	}, nil
}

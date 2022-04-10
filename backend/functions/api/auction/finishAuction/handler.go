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
	FinishAuction(ctx context.Context, auctionID string) error
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

	err := h.auctionService.FinishAuction(ctx, req.AuctionID)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNoContent,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}, nil
}

package main

import (
	"auctionsPlatform/utils"
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type request struct {
	ItemID string
}

type itemService interface {
	DeleteItem(ctx context.Context, itemID string) error
}

type handler struct {
	itemService itemService
}

// Reik pridet kad eitu submitint nuotraukas
func (h *handler) DeleteItem(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	accessToken := event.Headers["access_token"]
	if len(accessToken) <= 0 {
		return utils.InternalError("token not provided")
	}

	// userConfig, err := utils.GetUserConfig(accessToken)
	// if err != nil {
	// 	return utils.InternalError(err.Error())
	// }

	req := request{}

	err := h.itemService.DeleteItem(ctx, req.ItemID)
	if err != nil {
		return utils.InternalError(err.Error())
	}
	if err != nil {
		return utils.InternalError(err.Error())
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "*",
		},
		StatusCode: http.StatusNoContent,
	}, nil
}

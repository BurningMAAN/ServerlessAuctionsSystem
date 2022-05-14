package main

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type request struct {
	UserName    string
	ItemID      string
	Category    *string `json:"category"`
	Description *string `json:"description"`
	Name        *string `json:"name"`
}

type itemService interface {
	UpdateItem(ctx context.Context, itemID string, update models.ItemUpdate) error
}

type handler struct {
	itemService itemService
}

func (h *handler) SearchUserItems(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Print("got headers", event.Headers)
	accessToken := event.Headers["access_token"]
	if len(accessToken) <= 0 {
		return utils.InternalError("token not provided")
	}

	_, err := utils.GetUserConfig(accessToken)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	if len(event.PathParameters["userId"]) == 0 {
		return utils.InternalError("not provided auctionId")
	}

	if len(event.PathParameters["itemId"]) == 0 {
		return utils.InternalError("not provided auctionId")
	}

	req := request{
		UserName: event.PathParameters["userId"],
		ItemID:   event.PathParameters["itemId"],
	}

	log.Printf("req struct: %v", req)
	log.Printf("req body: %v", string(event.Body))
	err = json.Unmarshal([]byte(event.Body), &req)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	err = h.itemService.UpdateItem(ctx, req.ItemID, models.ItemUpdate{})
	if err != nil {
		return utils.InternalError(err.Error())
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: http.StatusOK,
	}, nil
}

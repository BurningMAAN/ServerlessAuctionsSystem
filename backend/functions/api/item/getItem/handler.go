package main

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type request struct {
	ItemID string
}

type response struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	OwnerID     string   `json:"ownerId"`
	PhotoURLs   []string `json:"photoURLs"`
}

type itemService interface {
	GetItemByID(ctx context.Context, itemID string) (models.Item, error)
}

type handler struct {
	itemService itemService
}

func (h *handler) GetItem(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if len(event.PathParameters["itemId"]) == 0 {
		return utils.InternalError("not provided auctionId")
	}
	req := request{
		ItemID: event.PathParameters["itemId"],
	}

	item, err := h.itemService.GetItemByID(ctx, req.ItemID)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	respBody, err := json.Marshal(response{
		ID:          item.ID,
		Description: item.Description,
		Category:    string(item.Category),
		OwnerID:     item.OwnerID,
		PhotoURLs:   item.PhotoURLs,
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

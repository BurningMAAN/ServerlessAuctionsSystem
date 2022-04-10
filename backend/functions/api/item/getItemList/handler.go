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
	// ItemID string
}

type response struct {
	Items []itemsResponse `json:"items"`
}

type itemsResponse struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	OwnerID     string   `json:"ownerId"`
	PhotoURLs   []string `json:"photoURLs"`
}

type itemService interface {
	GetItemsByUserID(ctx context.Context, userID string) ([]models.Item, error)
}

type handler struct {
	itemService itemService
}

func (h *handler) GetItems(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	items, err := h.itemService.GetItemsByUserID(ctx, "")
	if err != nil {
		return utils.InternalError(err.Error())
	}

	resp := response{}
	for _, item := range items {
		resp.Items = append(resp.Items, itemsResponse{
			ID:          item.ID,
			Description: item.Description,
			Category:    string(item.Category),
			OwnerID:     item.OwnerID,
			PhotoURLs:   item.PhotoURLs,
		})
	}
	respBody, err := json.Marshal(resp)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		StatusCode: http.StatusOK,
		Body:       string(respBody),
	}, nil
}

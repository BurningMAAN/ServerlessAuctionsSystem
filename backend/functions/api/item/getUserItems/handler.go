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
	UserID string
}

type response struct {
	ItemList []item `json:"items"`
}

type item struct {
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

func (h *handler) GetUserItems(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if len(event.PathParameters["userId"]) == 0 {
		return utils.InternalError("not provided auctionId")
	}
	req := request{
		UserID: event.PathParameters["userId"],
	}

	items, err := h.itemService.GetItemsByUserID(ctx, req.UserID)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	itemsList := []item{}
	for _, foundItem := range items {
		itemsList = append(itemsList, item{
			ID:          foundItem.ID,
			Description: foundItem.Description,
			Category:    string(foundItem.Category),
			OwnerID:     foundItem.OwnerID,
			PhotoURLs:   foundItem.PhotoURLs,
		})
	}

	respBody, err := json.Marshal(response{
		ItemList: itemsList})
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

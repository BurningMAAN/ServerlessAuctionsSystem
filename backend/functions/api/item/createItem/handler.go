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
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	OwnerID     string `json:"ownerId"`
}

type response struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	OwnerID     string   `json:"ownerId"`
	PhotoURLs   []string `json:"photoURLs"`
	Name        string   `json:"name"`
}

type itemService interface {
	CreateItem(ctx context.Context, item models.Item) (models.Item, error)
}

type handler struct {
	itemService itemService
}

// Reik pridet kad eitu submitint nuotraukas
func (h *handler) CreateItem(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req := request{}
	err := json.Unmarshal([]byte(event.Body), &req)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	item, err := h.itemService.CreateItem(ctx, models.Item{
		Description: req.Description,
		Category:    models.ItemCategory(req.Category),
		OwnerID:     req.OwnerID,
		PhotoURLs:   []string{},
		Name:        req.Name,
	})
	if err != nil {
		return utils.InternalError(err.Error())
	}

	respBody, err := json.Marshal(response{
		ID:          item.ID,
		Description: item.Description,
		Category:    string(item.Category),
		OwnerID:     item.OwnerID,
		PhotoURLs:   item.PhotoURLs,
		Name:        item.Name,
	})
	if err != nil {
		return utils.InternalError(err.Error())
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: http.StatusCreated,
		Body:       string(respBody),
	}, nil
}

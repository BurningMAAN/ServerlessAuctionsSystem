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
	UserName string
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
	GetItemsByUserName(ctx context.Context, userID string) ([]models.Item, error)
}

type handler struct {
	itemService itemService
}

func (h *handler) GetUserItems(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Print("got headers", event.Headers)
	accessToken := event.Headers["access_token"]
	if len(accessToken) <= 0 {
		return utils.InternalError("token not provided")
	}

	userConfig, err := utils.GetUserConfig(accessToken)
	if err != nil {
		return utils.InternalError(err.Error())
	}
	req := request{
		UserName: userConfig.Name,
	}

	items, err := h.itemService.GetItemsByUserName(ctx, req.UserName)
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

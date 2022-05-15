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
	Category   string `json:"category"`
	Type       string `json:"type"`
	WinnerName string `json:"winnerName"`
	Stage      string `json:"stage"`
}

type response struct {
	Auctions []auction `json:"auctions"`
}

type auction struct {
	ID           string       `json:"id"`
	AuctionDate  time.Time    `json:"auctionDate"`
	AuctionType  string       `json:"auctionType"`
	BidIncrement float64      `json:"bidIncrement"`
	IsFinished   bool         `json:"isFinished"`
	CreatorID    string       `json:"creatorId"`
	ItemID       string       `json:"itemId"`
	Item         itemResponse `json:"item"`
	PhotoURL     string       `json:"photoURL"`
}

type itemResponse struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	OwnerID     string   `json:"ownerId"`
	Name        string   `json:"name"`
	PhotoURLs   []string `json:"photoURLs"`
	AuctionID   string   `json:"auctionId"`
}

type auctionService interface {
	SearchAuctions(ctx context.Context, searchParams models.AuctionSearchParams) ([]models.AuctionListView, error)
}

type handler struct {
	auctionService auctionService
}

func (h *handler) GetAuction(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req := request{}

	err := json.Unmarshal([]byte(event.Body), &req)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	auctions, err := h.auctionService.SearchAuctions(ctx, models.AuctionSearchParams{
		Category:    &req.Category,
		AuctionType: &req.Type,
		WinnerName:  &req.WinnerName,
		Stage:       &req.Stage,
	})
	if err != nil {
		return utils.InternalError(err.Error())
	}

	resp := auctionsToResponse(auctions)
	respBody, err := json.Marshal(resp)
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

func auctionsToResponse(auctions []models.AuctionListView) response {
	auctionsList := []auction{}
	for _, auctionItem := range auctions {
		auctionsList = append(auctionsList, auction{
			ID:           auctionItem.Auction.ID,
			AuctionDate:  auctionItem.Auction.StartDate,
			AuctionType:  string(auctionItem.Auction.Type),
			BidIncrement: auctionItem.Auction.BidIncrement,
			CreatorID:    auctionItem.Auction.CreatorID,
			PhotoURL:     auctionItem.Auction.PhotoURL,
			Item: itemResponse{
				ID:          auctionItem.Auction.ItemID,
				Description: auctionItem.Item.Description,
				Category:    string(auctionItem.Item.Category),
				PhotoURLs:   auctionItem.Item.PhotoURLs,
				Name:        auctionItem.Item.Name,
				AuctionID:   auctionItem.Item.AuctionID,
			},
		})
	}

	return response{auctionsList}
}

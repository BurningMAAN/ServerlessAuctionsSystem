package main

import (
	"auctionsPlatform/models"
	bidRepo "auctionsPlatform/repositories/bid"
	"auctionsPlatform/utils"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type request struct {
	ItemID       string    `json:"itemId"`
	AuctionDate  time.Time `json:"auctionDate"`
	BuyoutPrice  *float64  `json:"buyoutPrice"`
	AuctionType  string    `json:"auctionType"`
	BidIncrement float64   `json:"bidIncrement"`
}

type response struct {
	Data map[string]auctionResponse
}

type auctionResponse struct {
	Auction models.Auction `json:"auction"`
	Item    models.Item    `json:"item"`
	Bids    []models.Bid   `json:"bids"`
}
type auctionService interface {
	GetAuctions(ctx context.Context) ([]models.AuctionListView, error)
}

type itemRepository interface {
	GetItemByID(ctx context.Context, itemID string) (models.Item, error)
}

type bidsRepository interface {
	GetLatestAuctionBids(ctx context.Context, auctionID string, optFns ...func(*bidRepo.OptionalGetParameters)) ([]models.Bid, error)
}

type handler struct {
	auctionService auctionService
	itemRepository itemRepository
	bidsRepository bidsRepository
}

func (h *handler) CreateAuction(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	auctions, err := h.auctionService.GetAuctions(ctx)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	response := response{}
	for _, auction := range auctions {
		item, err := h.itemRepository.GetItemByID(ctx, auction.Auction.ItemID)
		if err != nil {
			return utils.InternalError(err.Error())
		}

		bids, err := h.bidsRepository.GetLatestAuctionBids(ctx, auction.Auction.ID)
		if err != nil {
			return utils.InternalError(err.Error())
		}

		response.Data[auction.Auction.ID] = auctionResponse{
			Auction: auction.Auction,
			Item:    item,
			Bids:    bids,
		}
	}

	respBody, err := json.Marshal(response)
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

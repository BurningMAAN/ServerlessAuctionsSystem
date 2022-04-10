package auction

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func unmarshalAuction(auctionDB AuctionDB) (models.Auction, error) {
	return models.Auction{
		ID:           utils.Extract(models.AuctionEntityType, auctionDB.PK),
		Type:         models.AuctionType(auctionDB.Type),
		BuyoutPrice:  auctionDB.BuyoutPrice,
		BidIncrement: auctionDB.BidIncrement,
		StartDate:    auctionDB.StartDate,
		CreatorID:    auctionDB.CreatorID,
		IsFinished:   auctionDB.IsFinished,
		ItemID:       auctionDB.ItemID,
	}, nil
}

func ExtractAuction(items map[string]types.AttributeValue) (models.Auction, error) {
	auctionDB := AuctionDB{}
	err := attributevalue.UnmarshalMap(items, &auctionDB)
	if err != nil {
		return models.Auction{}, err
	}

	auction, err := unmarshalAuction(auctionDB)
	if err != nil {
		return models.Auction{}, err
	}

	return auction, nil
}

func ExtractAuctions(items []map[string]types.AttributeValue) ([]models.Auction, error) {
	auctions := []models.Auction{}
	for _, item := range items {
		auction, err := ExtractAuction(item)
		if err != nil {
			return []models.Auction{}, err
		}

		auctions = append(auctions, auction)
	}

	return auctions, nil
}

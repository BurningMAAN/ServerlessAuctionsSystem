package auction

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func marshalAuction(auctionID string, auction models.Auction) (AuctionDB, error) {
	return AuctionDB{
		PK:           fmt.Sprintf("Auction#%s", auctionID),
		SK:           "Metadata",
		BuyoutPrice:  auction.BuyoutPrice,
		StartDate:    auction.StartDate,
		BidIncrement: auction.BidIncrement,
		CreatorID:    auction.CreatorID,
	}, nil
}

func unmarshalAuction(auctionDB AuctionDB) (models.Auction, error) {
	return models.Auction{
		ID:           utils.Extract(models.AuctionEntityType, auctionDB.PK),
		Type:         models.AuctionType(auctionDB.Type),
		BuyoutPrice:  auctionDB.BuyoutPrice,
		BidIncrement: auctionDB.BidIncrement,
		StartDate:    auctionDB.StartDate,
		CreatorID:    auctionDB.CreatorID,
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

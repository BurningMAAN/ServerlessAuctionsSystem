package bid

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func unmarshalItem(bidDB BidDB) (models.Bid, error) {
	return models.Bid{
		ID:        utils.Extract(models.BidEntityType, bidDB.PK),
		Value:     bidDB.Value,
		Timestamp: bidDB.GSI1SK,
		AuctionID: utils.Extract(models.AuctionEntityType, bidDB.GSI1PK),
	}, nil
}

func ExtractBid(attrItem map[string]types.AttributeValue) (models.Bid, error) {
	itemDB := BidDB{}
	err := attributevalue.UnmarshalMap(attrItem, &itemDB)
	if err != nil {
		return models.Bid{}, err
	}

	item, err := unmarshalItem(itemDB)
	if err != nil {
		return models.Bid{}, err
	}

	return item, nil
}

func ExtractBids(items []map[string]types.AttributeValue) ([]models.Bid, error) {
	bids := []models.Bid{}
	for _, item := range items {
		bid, err := ExtractBid(item)
		if err != nil {
			return []models.Bid{}, err
		}

		bids = append(bids, bid)
	}

	return bids, nil
}

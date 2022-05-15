package auction

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func unmarshalAuction(auctionDB AuctionDB) (models.Auction, error) {
	return models.Auction{
		ID:           utils.Extract(models.AuctionEntityType, auctionDB.PK),
		Type:         models.AuctionType(auctionDB.Type),
		BidIncrement: auctionDB.BidIncrement,
		StartDate:    auctionDB.StartDate,
		CreatorID:    auctionDB.CreatorID,
		ItemID:       utils.Extract(models.ItemEntityType, auctionDB.GSI1PK),
		Category:     utils.Extract("Category", auctionDB.GSI1SK),
		EndDate:      auctionDB.EndDate,
		Stage:        auctionDB.Stage,
		PhotoURL:     auctionDB.PhotoURL,
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

func buildSearchCondition(searchParams models.AuctionSearchParams) expression.ConditionBuilder {
	conditionExpression := expression.ConditionBuilder{}
	if searchParams.Category != nil {
		conditionExpression = conditionExpression.And(expression.Name("GSI1SK").Equal(expression.Value(&searchParams.Category)))
	}
	if searchParams.AuctionType != nil {
		conditionExpression = conditionExpression.And(expression.Name("Type").Equal(expression.Value(&searchParams.AuctionType)))
	}

	return conditionExpression
}

func buildUpdate(update models.AuctionUpdate) expression.UpdateBuilder {
	updateExpression := expression.UpdateBuilder{}

	if update.BidIncrement != nil {
		updateExpression = updateExpression.Set(expression.Name("BidIncrement"), expression.Value(&update.BidIncrement))
	}

	if update.StartDate != nil {
		updateExpression = updateExpression.Set(expression.Name("StartDate"), expression.Value(&update.StartDate))
	}

	if update.Type != nil {
		updateExpression = updateExpression.Set(expression.Name("Type"), expression.Value(&update.Type))
	}

	if update.WinnerID != nil {
		updateExpression = updateExpression.Set(expression.Name("WinnerID"), expression.Value(&update.WinnerID))
	}

	return updateExpression
}

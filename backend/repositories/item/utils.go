package item

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"
	"log"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func unmarshalItem(itemDB ItemDB) (models.Item, error) {
	return models.Item{
		ID:          utils.Extract(models.ItemEntityType, itemDB.PK),
		Description: itemDB.Description,
		Category:    models.ItemCategory(itemDB.GSI1SK),
		OwnerID:     utils.Extract(models.UserEntityType, itemDB.GSI1PK),
		PhotoURLs:   itemDB.PhotoURLs,
		Name:        itemDB.Name,
		AuctionID:   utils.Extract(models.AuctionEntityType, itemDB.GSI2PK),
	}, nil
}

func ExtractItem(items map[string]types.AttributeValue) (models.Item, error) {
	itemDB := ItemDB{}
	err := attributevalue.UnmarshalMap(items, &itemDB)
	if err != nil {
		return models.Item{}, err
	}

	item, err := unmarshalItem(itemDB)
	if err != nil {
		return models.Item{}, err
	}

	return item, nil
}

func ExtractItems(items []map[string]types.AttributeValue) ([]models.Item, error) {
	returnedItems := []models.Item{}

	for _, item := range items {
		extractedItem, err := ExtractItem(item)
		if err != nil {
			return []models.Item{}, err
		}

		returnedItems = append(returnedItems, extractedItem)
	}

	return returnedItems, nil
}

func buildItemUpdate(update models.ItemUpdate) expression.UpdateBuilder {
	updateExpression := expression.UpdateBuilder{}
	if update.AuctionID != nil {
		updateExpression = updateExpression.Set(expression.Name("GSI2PK"), expression.Value(utils.Make(models.AuctionEntityType, *update.AuctionID)))
	}

	if update.OwnerID != nil {
		updateExpression = updateExpression.Set(expression.Name("GSI1PK"), expression.Value(utils.Make(models.UserEntityType, *update.OwnerID)))
	}

	if update.Category != nil {
		updateExpression = updateExpression.Set(expression.Name("GSI1SK"), expression.Value(utils.Make("Category", *update.Category)))
	}

	if update.Description != nil {
		updateExpression = updateExpression.Set(expression.Name("Description"), expression.Value(*update.Description))
	}

	if update.Name != nil {
		log.Print("I AM HERE")
		updateExpression = updateExpression.Set(expression.Name("Name"), expression.Value(&update.Name))
	}

	return updateExpression
}

func buildSearchFilter(searchParams models.ItemSearchParams) expression.ConditionBuilder {
	conditionExpression := expression.ConditionBuilder{}
	if searchParams.Category != nil {
		conditionExpression = conditionExpression.And(expression.Name("GSI1SK").Equal(expression.Value(&searchParams.Category)))
	}

	conditionExpression = conditionExpression.And(expression.Name("GSI1PK").Equal(expression.Value(utils.Make(models.UserEntityType, searchParams.OwnerID))))
	return conditionExpression
}

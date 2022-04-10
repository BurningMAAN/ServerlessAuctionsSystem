package item

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
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

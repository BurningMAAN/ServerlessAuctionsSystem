package item

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func marshalItem(itemID string, item models.Item) (ItemDB, error) {
	return ItemDB{
		PK:          fmt.Sprintf("Item#%s", itemID),
		SK:          "Metadata",
		GSI1PK:      utils.Make(models.UserEntityType, item.OwnerID),
		GSI1SK:      string(item.Category),
		PhotoURLs:   item.PhotoURLs,
		Description: item.Description,
	}, nil
}

func unmarshalItem(itemDB ItemDB) (models.Item, error) {
	return models.Item{
		ID:          utils.Extract(models.ItemEntityType, itemDB.PK),
		Description: itemDB.Description,
		Category:    models.ItemCategory(itemDB.GSI1SK),
		OwnerID:     utils.Extract(models.UserEntityType, itemDB.GSI1PK),
		PhotoURLs:   itemDB.PhotoURLs,
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

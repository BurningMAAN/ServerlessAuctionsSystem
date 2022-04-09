package item

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type DB interface {
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#Client.GetItem
	GetItem(ctx context.Context, input *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#Client.PutItem
	PutItem(ctx context.Context, input *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

type repository struct {
	tableName string
	DB        DB
}

func New(tableName string, db DB) *repository {
	return &repository{
		tableName: tableName,
		DB:        db,
	}
}

type ItemDB struct {
	PK          string // Example: Item#{ItemID}
	SK          string // Example: Metadata
	GSI1PK      string // Example: User#{OwnerID}
	GSI1SK      string // Example: Category
	PhotoURLs   []string
	Description string
}

func (r *repository) CreateItem(ctx context.Context, item models.Item) (models.Item, error) {
	itemID := uuid.New().String()

	itemDB := ItemDB{
		PK:          utils.Make(models.ItemEntityType, itemID),
		SK:          "Metadata",
		GSI1PK:      utils.Make(models.UserEntityType, item.OwnerID),
		GSI1SK:      string(item.Category),
		PhotoURLs:   item.PhotoURLs,
		Description: item.Description,
	}

	itemAttributeValues, err := attributevalue.MarshalMap(itemDB)
	if err != nil {
		return models.Item{}, err
	}

	query := &dynamodb.PutItemInput{
		Item:                itemAttributeValues,
		TableName:           aws.String(r.tableName),
		ConditionExpression: aws.String("attribute_not_exists(SK)"),
	}
	_, err = r.DB.PutItem(ctx, query)
	if err != nil {
		var ccfe *types.ConditionalCheckFailedException
		if errors.As(err, &ccfe) {
			return models.Item{}, errors.New("not exists")
		}
		return models.Item{}, err
	}

	item.ID = itemID
	return item, nil
}

func (r *repository) GetItemByID(ctx context.Context, itemID string) (models.Item, error) {
	query := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: utils.Make(models.ItemEntityType, itemID)},
			"SK": &types.AttributeValueMemberS{Value: "Metadata"},
		},
	}

	result, err := r.DB.GetItem(ctx, query)
	if err != nil {
		return models.Item{}, err
	}

	if result.Item == nil {
		return models.Item{}, errors.New("exists")
	}

	return ExtractItem(result.Item)
}

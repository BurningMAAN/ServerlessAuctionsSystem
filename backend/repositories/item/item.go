package item

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type DB interface {
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#Client.GetItem
	GetItem(ctx context.Context, input *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#Client.PutItem
	PutItem(ctx context.Context, input *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#Client.UpdateItem
	UpdateItem(ctx context.Context, input *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
	Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)
	Query(ctx context.Context, input *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
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
	PK     string // Example: Item#{ItemID}
	SK     string // Example: Metadata
	GSI1PK string `dynamodbav:",omitempty"` // Example: User#{OwnerID}
	GSI1SK string `dynamodbav:",omitempty"` // Example: Category#{Category}
	// GSI2PK      string `dynamodbav:",omitempty"` // Example: Auction#{AuctionID}
	// GSI2SK      string `dynamodbav:",omitempty"` // Example: Metadata
	PhotoURLs   []string
	Description string
	Name        string
	AuctionID   string
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
		Name:        item.Name,
		AuctionID:   item.AuctionID,
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

func (r *repository) AssignItem(ctx context.Context, auctionID, itemID string) error {
	express, err := expression.NewBuilder().WithUpdate(expression.Set(
		expression.Name("GSI2PK"), expression.Value(utils.Make(models.UserEntityType, auctionID)))).Build()
	if err != nil {
		return err
	}
	input := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(r.tableName),
		ReturnValues:              types.ReturnValueAllNew,
		ExpressionAttributeValues: express.Values(),
		ExpressionAttributeNames:  express.Names(),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{
				Value: fmt.Sprintf("%s#%s", models.ItemEntityType, itemID),
			},
			"SK": &types.AttributeValueMemberS{
				Value: "Metadata",
			},
		},
		UpdateExpression: express.Update(),
	}

	_, err = r.DB.UpdateItem(ctx, input)
	if err != nil {
		return err
	}

	return err
}

func (r *repository) GetItemsByUserName(ctx context.Context, userName string) ([]models.Item, error) {
	keyCondition := expression.Key("GSI1PK").Equal(expression.Value(utils.Make(models.UserEntityType, userName)))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		IndexName:                 aws.String("GSI1"),
		TableName:                 aws.String(r.tableName),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}

	res, err := r.DB.Query(ctx, queryInput)

	if err != nil {
		return nil, err
	}

	if len(res.Items) == 0 {
		return nil, nil
	}

	return ExtractItems(res.Items)
}

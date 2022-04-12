package bid

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"
	"context"
	"errors"
	"time"

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

type BidDB struct {
	PK     string // Example: Bid#{BidID}
	SK     string // Example: Metadata
	Value  float64
	GSI1PK string `dynamodbav:",omitempty"` // Example: Auction#{AuctionID}
	GSI1SK string `dynamodbav:",omitempty"` // Example: DateTime#2020-11-26T10:56:52Z
}

type OptionalGetParameters struct{}

func (r *repository) CreateBid(ctx context.Context, auctionID string, bid models.Bid) (models.Bid, error) {
	bidID := uuid.New().String()
	bidTime := time.Now()
	auctionDB := BidDB{
		PK:     utils.Make(models.BidEntityType, bidID),
		SK:     "Metadata",
		Value:  bid.Value,
		GSI1PK: utils.Make(models.AuctionEntityType, auctionID),
		GSI1SK: bidTime.String(),
	}

	bidAttributeValues, err := attributevalue.MarshalMap(auctionDB)
	if err != nil {
		return models.Bid{}, err
	}

	query := &dynamodb.PutItemInput{
		Item:                bidAttributeValues,
		TableName:           aws.String(r.tableName),
		ConditionExpression: aws.String("attribute_not_exists(SK)"),
	}
	_, err = r.DB.PutItem(ctx, query)
	if err != nil {
		var ccfe *types.ConditionalCheckFailedException
		if errors.As(err, &ccfe) {
			return models.Bid{}, errors.New("already exists")
		}
		return models.Bid{}, err
	}
	bid.ID = auctionID
	return bid, nil
}

func (r *repository) GetBidByID(ctx context.Context, bidID string) (models.Bid, error) {
	query := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: utils.Make(models.BidEntityType, bidID)},
			"SK": &types.AttributeValueMemberS{Value: "Metadata"},
		},
	}

	result, err := r.DB.GetItem(ctx, query)
	if err != nil {
		return models.Bid{}, err
	}

	if result.Item == nil {
		return models.Bid{}, errors.New("exists")
	}

	return ExtractBid(result.Item)
}

func (r *repository) GetLatestAuctionBids(ctx context.Context, auctionID string, optFns ...func(*OptionalGetParameters)) ([]models.Bid, error) {
	GSI1PK := utils.Make(models.AuctionEntityType, auctionID)
	keyCondition := expression.Key(string("GSI1PK")).Equal(expression.Value(GSI1PK))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()
	if err != nil {
		return nil, err
	}
	result, err := r.DB.Query(ctx, &dynamodb.QueryInput{
		TableName:                 &r.tableName,
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})
	if err != nil {
		return []models.Bid{}, err
	}

	if result.Items == nil {
		return []models.Bid{}, nil
	}

	return ExtractBids(result.Items)
}

package auction

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"
	"context"
	"errors"
	"time"

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

type AuctionDB struct {
	PK           string // Example: Auction#{AuctionID}
	SK           string // Example: Metadata
	BuyoutPrice  *float64
	StartDate    time.Time
	BidIncrement float64
	CreatorID    string
	Type         string
}

func (r *repository) CreateAuction(ctx context.Context, auction models.Auction) (models.Auction, error) {
	auctionID := uuid.New().String()

	auctionDB := AuctionDB{
		PK:           utils.Make(models.AuctionEntityType, auctionID),
		SK:           "Metadata",
		BuyoutPrice:  auction.BuyoutPrice,
		StartDate:    auction.StartDate,
		BidIncrement: auction.BidIncrement,
		CreatorID:    auction.CreatorID,
		Type:         string(auction.Type),
	}
	auctionAttributeValues, err := attributevalue.MarshalMap(auctionDB)
	if err != nil {
		return models.Auction{}, err
	}

	query := &dynamodb.PutItemInput{
		Item:                auctionAttributeValues,
		TableName:           aws.String(r.tableName),
		ConditionExpression: aws.String("attribute_not_exists(SK)"),
	}
	_, err = r.DB.PutItem(ctx, query)
	if err != nil {
		var ccfe *types.ConditionalCheckFailedException
		if errors.As(err, &ccfe) {
			return models.Auction{}, errors.New("not exists")
		}
		return models.Auction{}, err
	}
	auction.ID = auctionID
	return auction, nil
}

func (r *repository) GetAuctionByID(ctx context.Context, auctionID string) (models.Auction, error) {
	query := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: utils.Make(models.AuctionEntityType, auctionID)},
			"SK": &types.AttributeValueMemberS{Value: "Metadata"},
		},
	}

	result, err := r.DB.GetItem(ctx, query)
	if err != nil {
		return models.Auction{}, err
	}

	if result.Item == nil {
		return models.Auction{}, errors.New("exists")
	}

	return ExtractAuction(result.Item)
}

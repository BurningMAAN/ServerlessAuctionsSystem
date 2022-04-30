package auction

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"
	"context"
	"errors"
	"fmt"
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
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#Client.UpdateItem
	UpdateItem(ctx context.Context, input *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
	Query(ctx context.Context, input *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
	Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)
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
	PK             string // Example: Auction#{AuctionID}
	SK             string // Example: Metadata
	BuyoutPrice    *float64
	StartDate      time.Time
	BidIncrement   float64
	AuctionEndDate time.Time
	CreatorID      string
	Type           string
	IsFinished     bool
	ItemID         string
}

type OptionalGetParameters struct{}

func (r *repository) CreateAuction(ctx context.Context, auction models.Auction) (models.Auction, error) {
	auctionID := uuid.New().String()

	auctionDB := AuctionDB{
		PK:             utils.Make(models.AuctionEntityType, auctionID),
		SK:             "Metadata",
		BuyoutPrice:    auction.BuyoutPrice,
		StartDate:      auction.StartDate,
		BidIncrement:   auction.BidIncrement,
		CreatorID:      auction.CreatorID,
		Type:           string(auction.Type),
		ItemID:         auction.ItemID,
		IsFinished:     false,
		AuctionEndDate: auction.StartDate.Add(time.Duration(5 * time.Second)),
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

	err = r.CreateAuctionWorker(ctx, auction.ID, auction.StartDate, auctionDB.AuctionEndDate)
	if err != nil {

		return auction, err
	}

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
		return models.Auction{}, errors.New("resource does not exist")
	}

	return ExtractAuction(result.Item)
}

func (r *repository) FinishAuction(ctx context.Context, auctionID string) error {
	express, err := expression.NewBuilder().WithUpdate(expression.Set(
		expression.Name("IsFinished"), expression.Value(true))).Build()
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
				Value: fmt.Sprintf("%s#%s", models.AuctionEntityType, auctionID),
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

func (r *repository) GetAllAuctions(ctx context.Context, optFns ...func(*OptionalGetParameters)) ([]models.Auction, error) {
	filter := expression.Name("PK").BeginsWith(string(models.AuctionEntityType))

	expr, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		return []models.Auction{}, err
	}
	result, err := r.DB.Scan(ctx, &dynamodb.ScanInput{
		TableName:                 &r.tableName,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	})
	if err != nil {
		return []models.Auction{}, err
	}

	if result.Items == nil {
		return []models.Auction{}, errors.New("exists")
	}

	return ExtractAuctions(result.Items)
}

type auctionWorkerDB struct {
	PK        string
	SK        string
	StartDate time.Time
	EndDate   time.Time
}

func (r *repository) CreateAuctionWorker(ctx context.Context, auctionID string, startDate, endDate time.Time) error {
	auctionWorkerDB := auctionWorkerDB{
		PK:        utils.Make("AuctionWorker", auctionID),
		SK:        "Metadata",
		StartDate: startDate,
		EndDate:   endDate,
	}

	auctionAttributeValues, err := attributevalue.MarshalMap(auctionWorkerDB)
	if err != nil {
		return err
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
			return errors.New("not exists")
		}
		return err
	}
	return nil
}

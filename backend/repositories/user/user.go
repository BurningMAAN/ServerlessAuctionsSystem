package user

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
	"golang.org/x/crypto/bcrypt"
)

type UserDB struct {
	PK       string // Example: User#{UserID}
	SK       string // Example: Metadata
	GSI1PK   string `dynamodbav:",omitempty"` // Example: User#{UserName}
	GSI1SK   string `dynamodbav:",omitempty"` // Example: Metadata
	Password string
	Email    string
}

type DB interface {
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#Client.GetItem
	GetItem(ctx context.Context, input *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#Client.PutItem
	PutItem(ctx context.Context, input *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#Client.UpdateItem
	UpdateItem(ctx context.Context, input *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
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

func (r *repository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	userID := uuid.New().String()

	pswHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	itemDB := UserDB{
		PK:       utils.Make(models.UserEntityType, userID),
		SK:       "Metadata",
		GSI1PK:   utils.Make(models.UserEntityType, user.UserName),
		GSI1SK:   "Metadata",
		Password: string(pswHash),
		Email:    user.Email,
	}

	userAttributeValues, err := attributevalue.MarshalMap(itemDB)
	if err != nil {
		return models.User{}, err
	}

	query := &dynamodb.PutItemInput{
		Item:                userAttributeValues,
		TableName:           aws.String(r.tableName),
		ConditionExpression: aws.String("attribute_not_exists(SK)"),
	}
	_, err = r.DB.PutItem(ctx, query)
	if err != nil {
		var ccfe *types.ConditionalCheckFailedException
		if errors.As(err, &ccfe) {
			return models.User{}, errors.New("not exists")
		}
		return models.User{}, err
	}

	user.ID = userID
	return user, nil
}

func (r *repository) GetUserByID(ctx context.Context, userID string) (models.User, error) {
	query := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: utils.Make(models.UserEntityType, userID)},
			"SK": &types.AttributeValueMemberS{Value: "Metadata"},
		},
	}

	result, err := r.DB.GetItem(ctx, query)
	if err != nil {
		return models.User{}, err
	}

	if result.Item == nil {
		return models.User{}, errors.New("exists")
	}

	return ExtractUser(result.Item)
}

func (r *repository) GetUserByUserName(ctx context.Context, userName string) (models.User, error) {
	query := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"GSI1PK": &types.AttributeValueMemberS{Value: utils.Make(models.UserEntityType, userName)},
			"GSI1SK": &types.AttributeValueMemberS{Value: "Metadata"},
		},
	}

	result, err := r.DB.GetItem(ctx, query)
	if err != nil {
		return models.User{}, err
	}

	if result.Item == nil {
		return models.User{}, errors.New("exists")
	}

	return ExtractUser(result.Item)
}

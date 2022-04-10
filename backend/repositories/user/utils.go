package user

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func unmarshalUser(userDB UserDB) (models.User, error) {
	return models.User{
		UserName: utils.Extract(models.UserEntityType, userDB.PK),
		Password: userDB.Password,
		Email:    userDB.Email,
	}, nil
}

func ExtractUser(items map[string]types.AttributeValue) (models.User, error) {
	userDB := UserDB{}
	err := attributevalue.UnmarshalMap(items, &userDB)
	if err != nil {
		return models.User{}, err
	}

	user, err := unmarshalUser(userDB)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

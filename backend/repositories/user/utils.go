package user

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"
	"log"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"golang.org/x/crypto/bcrypt"
)

func unmarshalUser(userDB UserDB) (models.User, error) {
	return models.User{
		UserName: utils.Extract(models.UserEntityType, userDB.PK),
		Password: userDB.Password,
		Email:    userDB.Email,
		Role:     userDB.Role,
		Credit:   userDB.CreditBalance,
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

func buildUserUpdate(update models.UserUpdate) (expression.UpdateBuilder, error) {
	updateExpr := expression.UpdateBuilder{}
	log.Print("Cia atejo")
	if update.Password != nil {
		encryptedPsw, err := bcrypt.GenerateFromPassword([]byte(*update.Password), bcrypt.DefaultCost)
		if err != nil {
			return expression.UpdateBuilder{}, err
		}
		updateExpr = updateExpr.Set(expression.Name("Password"), expression.Value(string(encryptedPsw)))
	}
	if update.Email != nil {
		updateExpr = updateExpr.Set(expression.Name("Email"), expression.Value(string(*update.Email)))
	}
	if update.Credit != nil {
		log.Print("cia neatejo")
		updateExpr = updateExpr.Set(expression.Name("CreditBalance"), expression.Value(float64(*update.Credit)))
	}

	return updateExpr, nil
}

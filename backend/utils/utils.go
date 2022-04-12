package utils

import (
	"auctionsPlatform/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/dgrijalva/jwt-go"
)

type APIError struct {
	StatusCode   int    `json:"statusCode"`
	ErrorMessage string `json:"errorMessage"`
}

func InternalError(message string) (events.APIGatewayProxyResponse, error) {
	error := APIError{
		StatusCode:   http.StatusInternalServerError,
		ErrorMessage: message,
	}

	log.Print(message)
	errMsg, _ := json.Marshal(error)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "*",
		},
		Body: string(errMsg),
	}, nil
}

func Extract(entityType models.EntityType, key string) string {
	entity := fmt.Sprintf("%s#", entityType)
	extract := strings.TrimPrefix(key, entity)
	if extract == key {
		return ""
	}
	return extract
}

func Make(entityType models.EntityType, attrs ...string) string {
	if len(attrs) == 0 {
		return fmt.Sprintf("%s#", entityType)
	}

	return fmt.Sprintf("%s#%s", entityType, strings.Join(attrs, "#"))
}

func GetUserConfig(accessToken string) (models.UserConfig, error) {
	accessToken, err := strconv.Unquote(accessToken)
	if err != nil {
		return models.UserConfig{}, err
	}
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})
	if err != nil {
		return models.UserConfig{}, fmt.Errorf("failed to parse token: %w", err)
	}

	userConfig := models.UserConfig{}
	for key, val := range claims {
		if key == "username" {
			userConfig.Name = val.(string)
		}
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}

	fmt.Println(token)

	return userConfig, nil
}

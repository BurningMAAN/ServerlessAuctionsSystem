package utils

import (
	"auctionsPlatform/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type APIError struct {
	StatusCode   int    `json:"statusCode"`
	ErrorMessage string `json:"errorMessage"`
}

func InternalError() (events.APIGatewayProxyResponse, error) {
	error := APIError{
		StatusCode:   http.StatusInternalServerError,
		ErrorMessage: "Internal Server Error",
	}

	errMsg, _ := json.Marshal(error)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Headers: map[string]string{
			"Content-Type": "application/json",
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

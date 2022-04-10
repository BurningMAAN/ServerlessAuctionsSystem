package main

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type response struct {
	Token string `json:"access_token"`
}
type authorizerService interface {
	Authorize(ctx context.Context, userName, password string) (models.AuthorizationConfig, error)
}

type handler struct {
	authorizerService authorizerService
}

func (h *handler) Authorize(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req := request{}
	err := json.Unmarshal([]byte(event.Body), &req)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	authorizationConfig, err := h.authorizerService.Authorize(ctx, req.Username, req.Password)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	respBody, err := json.Marshal(response{
		Token: authorizationConfig.Token,
	})
	if err != nil {
		return utils.InternalError(err.Error())
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: http.StatusOK,
		Body:       string(respBody),
	}, nil
}

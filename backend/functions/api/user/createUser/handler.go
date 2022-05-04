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
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type response struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type userService interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
}

type handler struct {
	userService userService
}

func (h *handler) GetItems(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req := request{}
	err := json.Unmarshal([]byte(event.Body), &req)
	if err != nil {
		return utils.InternalError("failed to unmarshal")
	}

	user, err := h.userService.CreateUser(ctx, models.User{
		UserName: req.UserName,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		return utils.InternalError(err.Error())
	}

	respBody, err := json.Marshal(response{
		ID:       user.ID,
		Username: user.UserName,
		Email:    user.Email,
	})
	if err != nil {
		return utils.InternalError(err.Error())
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: http.StatusCreated,
		Body:       string(respBody),
	}, nil
}

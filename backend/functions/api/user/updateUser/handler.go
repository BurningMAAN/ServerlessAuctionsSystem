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
	UserID   string
	Password string `json:"password"`
	Email    string `json:"email"`
}

type userService interface {
	UpdateUser(ctx context.Context, updateModel models.UserUpdate) error
}

type handler struct {
	userService userService
}

func (h *handler) UpdateUser(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if len(event.PathParameters["userId"]) == 0 {
		return utils.InternalError("not provided userId")
	}

	req := request{
		UserID: event.PathParameters["itemId"],
	}

	err := json.Unmarshal([]byte(event.Body), &req)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	updateModel := models.UserUpdate{
		ID:       req.UserID,
		Password: &req.Password,
		Email:    &req.Email,
	}

	err = h.userService.UpdateUser(ctx, updateModel)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	if err != nil {
		return utils.InternalError(err.Error())
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		StatusCode: http.StatusOK,
	}, nil
}

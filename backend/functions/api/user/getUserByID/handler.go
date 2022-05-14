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
	UserID string
}

type response struct {
	ID            string  `json:"id"`
	Username      string  `json:"username"`
	Email         string  `json:"email"`
	CreditBalance float64 `json:"creditBalance"`
}

type userService interface {
	GetUserByUserName(ctx context.Context, userName string) (models.User, error)
}

type handler struct {
	userService userService
}

func (h *handler) GetItems(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if len(event.PathParameters["userId"]) == 0 {
		return utils.InternalError("not provided userId")
	}

	req := request{
		UserID: event.PathParameters["userId"],
	}

	user, err := h.userService.GetUserByUserName(ctx, req.UserID)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	respBody, err := json.Marshal(response{
		ID:            user.ID,
		Username:      user.UserName,
		Email:         user.Email,
		CreditBalance: user.Credit,
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

package main

import (
	"auctionsPlatform/models"
	"auctionsPlatform/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type request struct {
	UserName      string   `json:"userId"`
	Password      *string  `json:"password"`
	Email         *string  `json:"email"`
	CreditBalance *float64 `json:"creditBalance"`
}

type userService interface {
	UpdateUser(ctx context.Context, updateModel models.UserUpdate) error
}

type handler struct {
	userService userService
}

func (h *handler) UpdateUser(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req := request{}
	err := json.Unmarshal([]byte(event.Body), &req)
	if err != nil {
		return utils.InternalError(err.Error())
	}

	log.Printf("req body: %s", event.Body)
	updateModel := models.UserUpdate{
		UserName: req.UserName,
		Password: req.Password,
		Email:    req.Email,
		Credit:   req.CreditBalance,
	}

	log.Printf("after unmarshal: %v", updateModel)
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

package main

import (
	"auctionsPlatform/models"
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

type MockItemService struct {
	createItem func(ctx context.Context, item models.Item) (models.Item, error)
}

func (m *MockItemService) CreateItem(ctx context.Context, item models.Item) (models.Item, error) {
	return m.createItem(ctx, item)
}

func Test_CreateItem(t *testing.T) {
	tests := []struct {
		name           string
		itemService    itemService
		event          events.APIGatewayProxyRequest
		expectedResult events.APIGatewayProxyResponse
		expectedErr    error
	}{
		{
			name: "successful func call",
			event: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3RpbmciLCJleHAiOjE2NDk2OTQ0MDJ9.CVXpTJSv90EW3OsdAs4sKL8-5S4A_w8LpjwVJcI1Evg",
				},
				Body: `{
					"id": "38585d69-bd40-4909-8e4e-4873a3723e75",
					"description": "Best dvirka in za world",
					"category": "Transportas",
					"ownerId": "",
					"photoURLs": [],
					"name": "Dvirka"
				}`,
			},
			itemService: &MockItemService{
				createItem: func(ctx context.Context, item models.Item) (models.Item, error) {
					return models.Item{}, nil
				},
			},
		},
	}

	ctx := context.Background()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := handler{
				itemService: &MockItemService{},
			}

			res, err := h.CreateItem(ctx, test.event)
			assert.Equal(t, test.expectedResult, res)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}
func Test_tokenDecode(t *testing.T) {
	tests := []struct {
		name           string
		token          string
		expectedOutput UserConfig
		expectedErr    error
	}{
		{
			name:           "successful func call",
			token:          "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3RpbmciLCJleHAiOjE2NDk2OTQ0MDJ9.CVXpTJSv90EW3OsdAs4sKL8-5S4A_w8LpjwVJcI1Evg",
			expectedOutput: UserConfig{},
			expectedErr:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tokenConfig, err := getUserConfig(test.token)
			assert.Equal(t, test.expectedOutput, tokenConfig)
			assert.Equal(t, test.expectedErr, err)
			t.Fail()
		})
	}
}

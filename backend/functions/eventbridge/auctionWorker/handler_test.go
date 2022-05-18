package main

// import (
// 	"auctionsPlatform/models"
// 	bidRepo "auctionsPlatform/repositories/bid"
// 	"context"
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// )

// type mockAuctionRepository struct {
// 	updateAuctionStage   func(ctx context.Context, auctionID string, stage string) error
// 	updateAuctionEndDate func(ctx context.Context, auctionID string, endDate time.Time) error
// 	updateAuction        func(ctx context.Context, auctionID string, update models.AuctionUpdate) error
// }

// func (m mockAuctionRepository) UpdateAuctionStage(ctx context.Context, auctionID string, stage string) error {
// 	return m.updateAuctionStage(ctx, auctionID, stage)
// }

// func (m mockAuctionRepository) UpdateAuctionEndDate(ctx context.Context, auctionID string, endDate time.Time) error {
// 	return m.updateAuctionEndDate(ctx, auctionID, endDate)
// }

// func (m mockAuctionRepository) UpdateAuction(ctx context.Context, auctionID string, update models.AuctionUpdate) error {
// 	return m.updateAuction(ctx, auctionID, update)
// }

// type mockBidRepository struct {
// 	getLatestAuctionBids func(ctx context.Context, auctionID string, optFns ...func(*bidRepo.OptionalGetParameters)) ([]models.Bid, error)
// }

// func (m mockBidRepository) GetLatestAuctionBids(ctx context.Context, auctionID string, optFns ...func(*bidRepo.OptionalGetParameters)) ([]models.Bid, error) {
// 	return m.getLatestAuctionBids(ctx, auctionID)
// }

// type mockEventRepository struct {
// 	updateEventRule func(ctx context.Context, auctionID string, newDate time.Time) error
// 	deleteEventRule func(ctx context.Context, auctionID string) error
// }

// func (m mockEventRepository) UpdateEventRule(ctx context.Context, auctionID string, newDate time.Time) error {
// 	return m.updateEventRule(ctx, auctionID, newDate)
// }

// func (m mockEventRepository) DeleteEventRule(ctx context.Context, auctionID string) error {
// 	return m.deleteEventRule(ctx, auctionID)
// }

// type mockUserRepository struct {
// 	updateUser        func(ctx context.Context, updateModel models.UserUpdate) error
// 	getUserByUserName func(ctx context.Context, userName string) (models.User, error)
// }

// func (m mockUserRepository) UpdateUser(ctx context.Context, updateModel models.UserUpdate) error {
// 	return m.updateUser(ctx, updateModel)
// }

// func (m mockUserRepository) GetUserByUserName(ctx context.Context, userName string) (models.User, error) {
// 	return m.getUserByUserName(ctx, userName)
// }

// func Test_HandleAuction(t *testing.T) {
// 	mockTime, err := time.Parse(time.RFC3339Nano, "2021-11-11T17:11:33.515292649Z")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	tests := []struct {
// 		name        string
// 		handler     handler
// 		event       models.AuctionEvent
// 		expectedErr error
// 	}{
// 		{
// 			name: "got event to start auction - successful",
// 			event: models.AuctionEvent{
// 				AuctionID: "auctionID",
// 				Stage:     "STAGE_ACCEPTING_BIDS",
// 				EndDate:   mockTime,
// 			},
// 			expectedErr: nil,
// 			handler: handler{
// 				auctionRepo: mockAuctionRepository{
// 					updateAuctionStage: func(ctx context.Context, auctionID string, stage string) error {
// 						return nil
// 					},
// 					updateAuctionEndDate: func(ctx context.Context, auctionID string, endDate time.Time) error {
// 						return nil
// 					},
// 				},
// 				eventRepository: mockEventRepository{
// 					updateEventRule: func(ctx context.Context, auctionID string, newDate time.Time) error {
// 						return nil
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "got event to start auction - update auction stage fails",
// 			event: models.AuctionEvent{
// 				AuctionID: "auctionID",
// 				Stage:     "STAGE_ACCEPTING_BIDS",
// 				EndDate:   mockTime,
// 			},
// 			expectedErr: errors.New("internal"),
// 			handler: handler{
// 				auctionRepo: mockAuctionRepository{
// 					updateAuctionStage: func(ctx context.Context, auctionID string, stage string) error {
// 						return errors.New("internal")
// 					},
// 					updateAuctionEndDate: func(ctx context.Context, auctionID string, endDate time.Time) error {
// 						return nil
// 					},
// 				},
// 				eventRepository: mockEventRepository{
// 					updateEventRule: func(ctx context.Context, auctionID string, newDate time.Time) error {
// 						return nil
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "got event to start auction - update auction end date fails",
// 			event: models.AuctionEvent{
// 				AuctionID: "auctionID",
// 				Stage:     "STAGE_ACCEPTING_BIDS",
// 				EndDate:   mockTime,
// 			},
// 			expectedErr: errors.New("internal"),
// 			handler: handler{
// 				auctionRepo: mockAuctionRepository{
// 					updateAuctionStage: func(ctx context.Context, auctionID string, stage string) error {
// 						return nil
// 					},
// 					updateAuctionEndDate: func(ctx context.Context, auctionID string, endDate time.Time) error {
// 						return errors.New("internal")
// 					},
// 				},
// 				eventRepository: mockEventRepository{
// 					updateEventRule: func(ctx context.Context, auctionID string, newDate time.Time) error {
// 						return nil
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "got event to start auction - update auction event rule fails",
// 			event: models.AuctionEvent{
// 				AuctionID: "auctionID",
// 				Stage:     "STAGE_ACCEPTING_BIDS",
// 				EndDate:   mockTime,
// 			},
// 			expectedErr: errors.New("internal"),
// 			handler: handler{
// 				auctionRepo: mockAuctionRepository{
// 					updateAuctionStage: func(ctx context.Context, auctionID string, stage string) error {
// 						return nil
// 					},
// 					updateAuctionEndDate: func(ctx context.Context, auctionID string, endDate time.Time) error {
// 						return nil
// 					},
// 				},
// 				eventRepository: mockEventRepository{
// 					updateEventRule: func(ctx context.Context, auctionID string, newDate time.Time) error {
// 						return errors.New("internal")
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "got event to end auction - successful",
// 			event: models.AuctionEvent{
// 				AuctionID: "auctionID",
// 				Stage:     "STAGE_AUCTION_ONGOING",
// 				EndDate:   mockTime,
// 			},
// 			expectedErr: nil,
// 			handler: handler{
// 				auctionRepo: mockAuctionRepository{
// 					updateAuctionStage: func(ctx context.Context, auctionID string, stage string) error {
// 						return nil
// 					},
// 					updateAuction: func(ctx context.Context, auctionID string, update models.AuctionUpdate) error {
// 						return nil
// 					},
// 				},
// 				eventRepository: mockEventRepository{
// 					deleteEventRule: func(ctx context.Context, auctionID string) error {
// 						return nil
// 					},
// 				},
// 				bidRepository: mockBidRepository{
// 					getLatestAuctionBids: func(ctx context.Context, auctionID string, optFns ...func(*bidRepo.OptionalGetParameters)) ([]models.Bid, error) {
// 						return []models.Bid{
// 							{
// 								ID:        "BidID",
// 								Value:     555.5,
// 								UserID:    "testUser",
// 								AuctionID: "AuctionID",
// 								Timestamp: mockTime,
// 							},
// 						}, nil
// 					},
// 				},
// 				userRepository: mockUserRepository{
// 					getUserByUserName: func(ctx context.Context, userName string) (models.User, error) {
// 						return models.User{
// 							UserName: "testUser",
// 							Credit:   666.5,
// 						}, nil
// 					},
// 					updateUser: func(ctx context.Context, updateModel models.UserUpdate) error {
// 						return nil
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "got event to end auction - fail to update auction stage",
// 			event: models.AuctionEvent{
// 				AuctionID: "auctionID",
// 				Stage:     "STAGE_AUCTION_ONGOING",
// 				EndDate:   mockTime,
// 			},
// 			expectedErr: errors.New("internal"),
// 			handler: handler{
// 				auctionRepo: mockAuctionRepository{
// 					updateAuctionStage: func(ctx context.Context, auctionID string, stage string) error {
// 						return errors.New("internal")
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "got event to end auction - delete event rule unsuccessful",
// 			event: models.AuctionEvent{
// 				AuctionID: "auctionID",
// 				Stage:     "STAGE_AUCTION_ONGOING",
// 				EndDate:   mockTime,
// 			},
// 			expectedErr: errors.New("internal"),
// 			handler: handler{
// 				auctionRepo: mockAuctionRepository{
// 					updateAuctionStage: func(ctx context.Context, auctionID string, stage string) error {
// 						return nil
// 					},
// 				},
// 				eventRepository: mockEventRepository{
// 					deleteEventRule: func(ctx context.Context, auctionID string) error {
// 						return errors.New("internal")
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "got event to end auction - failing to get latest bids",
// 			event: models.AuctionEvent{
// 				AuctionID: "auctionID",
// 				Stage:     "STAGE_AUCTION_ONGOING",
// 				EndDate:   mockTime,
// 			},
// 			expectedErr: errors.New("internal"),
// 			handler: handler{
// 				auctionRepo: mockAuctionRepository{
// 					updateAuctionStage: func(ctx context.Context, auctionID string, stage string) error {
// 						return nil
// 					},
// 				},
// 				eventRepository: mockEventRepository{
// 					deleteEventRule: func(ctx context.Context, auctionID string) error {
// 						return nil
// 					},
// 				},
// 				bidRepository: mockBidRepository{
// 					getLatestAuctionBids: func(ctx context.Context, auctionID string, optFns ...func(*bidRepo.OptionalGetParameters)) ([]models.Bid, error) {
// 						return []models.Bid{}, errors.New("internal")
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "got event to end auction - failing at getting user",
// 			event: models.AuctionEvent{
// 				AuctionID: "auctionID",
// 				Stage:     "STAGE_AUCTION_ONGOING",
// 				EndDate:   mockTime,
// 			},
// 			expectedErr: errors.New("internal"),
// 			handler: handler{
// 				auctionRepo: mockAuctionRepository{
// 					updateAuctionStage: func(ctx context.Context, auctionID string, stage string) error {
// 						return nil
// 					},
// 				},
// 				eventRepository: mockEventRepository{
// 					deleteEventRule: func(ctx context.Context, auctionID string) error {
// 						return nil
// 					},
// 				},
// 				bidRepository: mockBidRepository{
// 					getLatestAuctionBids: func(ctx context.Context, auctionID string, optFns ...func(*bidRepo.OptionalGetParameters)) ([]models.Bid, error) {
// 						return []models.Bid{
// 							{
// 								ID:        "BidID",
// 								Value:     555.5,
// 								UserID:    "testUser",
// 								AuctionID: "AuctionID",
// 								Timestamp: mockTime,
// 							},
// 						}, nil
// 					},
// 				},
// 				userRepository: mockUserRepository{
// 					getUserByUserName: func(ctx context.Context, userName string) (models.User, error) {
// 						return models.User{}, errors.New("internal")
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "got event to end auction - update auction winner fails",
// 			event: models.AuctionEvent{
// 				AuctionID: "auctionID",
// 				Stage:     "STAGE_AUCTION_ONGOING",
// 				EndDate:   mockTime,
// 			},
// 			expectedErr: errors.New("internal"),
// 			handler: handler{
// 				auctionRepo: mockAuctionRepository{
// 					updateAuctionStage: func(ctx context.Context, auctionID string, stage string) error {
// 						return nil
// 					},
// 					updateAuction: func(ctx context.Context, auctionID string, update models.AuctionUpdate) error {
// 						return errors.New("internal")
// 					},
// 				},
// 				eventRepository: mockEventRepository{
// 					deleteEventRule: func(ctx context.Context, auctionID string) error {
// 						return nil
// 					},
// 				},
// 				bidRepository: mockBidRepository{
// 					getLatestAuctionBids: func(ctx context.Context, auctionID string, optFns ...func(*bidRepo.OptionalGetParameters)) ([]models.Bid, error) {
// 						return []models.Bid{
// 							{
// 								ID:        "BidID",
// 								Value:     555.5,
// 								UserID:    "testUser",
// 								AuctionID: "AuctionID",
// 								Timestamp: mockTime,
// 							},
// 						}, nil
// 					},
// 				},
// 				userRepository: mockUserRepository{
// 					getUserByUserName: func(ctx context.Context, userName string) (models.User, error) {
// 						return models.User{
// 							UserName: "testUser",
// 							Credit:   666.5,
// 						}, nil
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "got event to end auction - failing to adjust user credit balance",
// 			event: models.AuctionEvent{
// 				AuctionID: "auctionID",
// 				Stage:     "STAGE_AUCTION_ONGOING",
// 				EndDate:   mockTime,
// 			},
// 			expectedErr: errors.New("internal"),
// 			handler: handler{
// 				auctionRepo: mockAuctionRepository{
// 					updateAuctionStage: func(ctx context.Context, auctionID string, stage string) error {
// 						return nil
// 					},
// 					updateAuction: func(ctx context.Context, auctionID string, update models.AuctionUpdate) error {
// 						return nil
// 					},
// 				},
// 				eventRepository: mockEventRepository{
// 					deleteEventRule: func(ctx context.Context, auctionID string) error {
// 						return nil
// 					},
// 				},
// 				bidRepository: mockBidRepository{
// 					getLatestAuctionBids: func(ctx context.Context, auctionID string, optFns ...func(*bidRepo.OptionalGetParameters)) ([]models.Bid, error) {
// 						return []models.Bid{
// 							{
// 								ID:        "BidID",
// 								Value:     555.5,
// 								UserID:    "testUser",
// 								AuctionID: "AuctionID",
// 								Timestamp: mockTime,
// 							},
// 						}, nil
// 					},
// 				},
// 				userRepository: mockUserRepository{
// 					getUserByUserName: func(ctx context.Context, userName string) (models.User, error) {
// 						return models.User{
// 							UserName: "testUser",
// 							Credit:   666.5,
// 						}, nil
// 					},
// 					updateUser: func(ctx context.Context, updateModel models.UserUpdate) error {
// 						return errors.New("internal")
// 					},
// 				},
// 			},
// 		},
// 	}

// 	ctx := context.Background()
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			err := test.handler.HandleAuction(ctx, test.event)
// 			assert.Equal(t, test.expectedErr, err)
// 		})
// 	}
// }

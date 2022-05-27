package integration

import (
	"context"
	"flag"
	"os"
	"testing"
	"time"

	"auctionsPlatform/errors"
	"auctionsPlatform/models"
	auctionRepository "auctionsPlatform/repositories/auction"
	bidRepository "auctionsPlatform/repositories/bid"
	"auctionsPlatform/repositories/eventbridge"
	itemRepository "auctionsPlatform/repositories/item"
	userRepository "auctionsPlatform/repositories/user"
	"auctionsPlatform/services/auction"
	"auctionsPlatform/services/bid"
	"auctionsPlatform/services/item"
	"auctionsPlatform/services/user"
	localdb "auctionsPlatform/tests/dynamodb"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/stretchr/testify/assert"
)

var testDB *localdb.DB

func TestMain(m *testing.M) {
	// flags need to be parsed manually if `testing.Short()` is called in TestMain
	flag.Parse()

	if testing.Short() {
		os.Exit(m.Run())
	}

	testDB = localdb.MustNew(localdb.SimpleTable())
	exitCode := m.Run()
	testDB.MustStop()
	os.Exit(exitCode)
}

type mockEventBridge struct {
	putRule       func(ctx context.Context, params *cloudwatchevents.PutRuleInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.PutRuleOutput, error)
	putTargets    func(ctx context.Context, params *cloudwatchevents.PutTargetsInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.PutTargetsOutput, error)
	deleteRule    func(ctx context.Context, params *cloudwatchevents.DeleteRuleInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.DeleteRuleOutput, error)
	removeTargets func(ctx context.Context, params *cloudwatchevents.RemoveTargetsInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.RemoveTargetsOutput, error)
	putEvents     func(ctx context.Context, params *cloudwatchevents.PutEventsInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.PutEventsOutput, error)
}

func (m mockEventBridge) PutRule(ctx context.Context, params *cloudwatchevents.PutRuleInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.PutRuleOutput, error) {
	return m.putRule(ctx, params)
}

func (m mockEventBridge) PutTargets(ctx context.Context, params *cloudwatchevents.PutTargetsInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.PutTargetsOutput, error) {
	return m.putTargets(ctx, params)
}

func (m mockEventBridge) DeleteRule(ctx context.Context, params *cloudwatchevents.DeleteRuleInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.DeleteRuleOutput, error) {
	return m.deleteRule(ctx, params)
}

func (m mockEventBridge) RemoveTargets(ctx context.Context, params *cloudwatchevents.RemoveTargetsInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.RemoveTargetsOutput, error) {
	return m.removeTargets(ctx, params)
}

func (m mockEventBridge) PutEvents(ctx context.Context, params *cloudwatchevents.PutEventsInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.PutEventsOutput, error) {
	return m.putEvents(ctx, params)
}

func Test_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("not running database integration tests when -short flag provided")
	}

	ctx := context.Background()
	// Initialize all repositories
	auctionRepo := auctionRepository.New("mockTable", testDB.DB)
	itemRepo := itemRepository.New("mockTable", testDB.DB)
	bidRepo := bidRepository.New("mockTable", testDB.DB)
	userRepo := userRepository.New("mockTable", testDB.DB)

	// Initialize all services
	itemService := item.New(itemRepo, auctionRepo)
	auctionService := auction.New(auctionRepo, itemRepo, eventbridge.New(mockEventBridge{
		putRule: func(ctx context.Context, params *cloudwatchevents.PutRuleInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.PutRuleOutput, error) {
			return &cloudwatchevents.PutRuleOutput{}, nil
		},
		putTargets: func(ctx context.Context, params *cloudwatchevents.PutTargetsInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.PutTargetsOutput, error) {
			return &cloudwatchevents.PutTargetsOutput{}, nil
		},
		deleteRule: func(ctx context.Context, params *cloudwatchevents.DeleteRuleInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.DeleteRuleOutput, error) {
			return &cloudwatchevents.DeleteRuleOutput{}, nil
		},
		removeTargets: func(ctx context.Context, params *cloudwatchevents.RemoveTargetsInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.RemoveTargetsOutput, error) {
			return &cloudwatchevents.RemoveTargetsOutput{}, nil
		},
		putEvents: func(ctx context.Context, params *cloudwatchevents.PutEventsInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.PutEventsOutput, error) {
			return &cloudwatchevents.PutEventsOutput{}, nil
		},
	}))

	userService := user.New(userRepo)
	bidService := bid.New(auctionRepo, bidRepo, userRepo)
	// Create Users
	user, err := userService.CreateUser(ctx, models.User{
		UserName: "testUser",
		Password: "testPassword",
		Email:    "testEmail@email.com",
		Role:     "standard_user",
		Credit:   0,
	})
	assert.Nil(t, err)

	bidder, err := userService.CreateUser(ctx, models.User{
		UserName: "testUser1",
		Password: "testPassword",
		Email:    "testEmail@email.com",
		Role:     "standard_user",
		Credit:   1000,
	})
	assert.Nil(t, err)

	// Check if created user is OK
	assert.Nil(t, err)
	assert.Equal(t, models.User{
		ID:       user.ID,
		UserName: "testUser",
		Password: "testPassword",
		Email:    "testEmail@email.com",
		Role:     "standard_user",
		Credit:   0,
	}, user)

	// Update user credit balance
	updateCredit := float64(1000)
	err = userService.UpdateUser(ctx, models.UserUpdate{
		UserName: "testUser",
		Credit:   &updateCredit,
	})
	assert.Nil(t, err)

	// Validate user credit balance
	gotUser, err := userService.GetUserByID(ctx, user.UserName)
	assert.Nil(t, err)
	assert.Equal(t, updateCredit, gotUser.Credit)

	// Create Item
	createdItem, err := itemService.CreateItem(ctx, models.Item{
		Name:        "Bike",
		Description: "very nice bike",
		Category:    models.ItemCategoryTransport,
		PhotoURLs:   []string{"example"},
		OwnerID:     user.UserName,
	})
	assert.Nil(t, err)
	assert.Equal(t, models.Item{
		ID:          createdItem.ID,
		Name:        "Bike",
		Description: "very nice bike",
		Category:    models.ItemCategoryTransport,
		PhotoURLs:   []string{"example"},
		OwnerID:     user.UserName,
	}, createdItem)

	// Create Auction for created item
	bidIncrement := float64(500)
	exampleStartDate, _ := time.Parse(time.RFC3339, "2023-10-02T15:00:00Z")
	exampleEndDate, _ := time.Parse(time.RFC3339, "2024-10-02T15:00:00Z")
	createdAuction, err := auctionService.CreateAuction(ctx, models.Auction{
		Type:         models.AuctionTypeAbsolute,
		StartDate:    exampleStartDate,
		EndDate:      exampleEndDate,
		BidIncrement: bidIncrement,
		CreatorID:    user.UserName,
	}, createdItem.ID)
	assert.Nil(t, err)

	// Get auction
	gotAuction, err := auctionService.GetAuctionByID(ctx, createdAuction.ID)
	assert.Nil(t, err)
	assert.Equal(t, createdAuction.ID, gotAuction.ID)

	firstBid, err := bidService.PlaceBid(ctx, createdAuction.ID, models.Bid{
		UserID:    bidder.UserName,
		Value:     500,
		AuctionID: createdAuction.ID,
		Timestamp: time.Now(),
	})
	assert.Nil(t, err)
	_, err = bidService.GetBidByID(ctx, firstBid.ID)
	assert.Nil(t, err)

	secondBid, err := bidService.PlaceBid(ctx, createdAuction.ID, models.Bid{
		UserID: bidder.UserName,
		Value:  500,
	})
	assert.Nil(t, err)
	gotSecondBid, err := bidService.GetBidByID(ctx, secondBid.ID)
	assert.Nil(t, err)
	assert.Equal(t, secondBid, gotSecondBid)

	_, err = bidService.PlaceBid(ctx, createdAuction.ID, models.Bid{
		UserID: bidder.UserName,
		Value:  500,
	})
	assert.Equal(t, errors.ErrUnsufficientCreditBalance, err)

	bids, err := bidService.GetLatestAuctionBids(ctx, createdAuction.ID)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(bids))
}

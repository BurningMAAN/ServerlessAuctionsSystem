package eventbridge

import (
	"auctionsPlatform/models"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	cloudwatchTypes "github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
	"github.com/aws/aws-sdk-go/aws"
)

type eventClient interface {
	PutRule(ctx context.Context, params *cloudwatchevents.PutRuleInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.PutRuleOutput, error)
	PutTargets(ctx context.Context, params *cloudwatchevents.PutTargetsInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.PutTargetsOutput, error)
	DeleteRule(ctx context.Context, params *cloudwatchevents.DeleteRuleInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.DeleteRuleOutput, error)
	RemoveTargets(ctx context.Context, params *cloudwatchevents.RemoveTargetsInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.RemoveTargetsOutput, error)
	PutEvents(ctx context.Context, params *cloudwatchevents.PutEventsInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.PutEventsOutput, error)
}

type repository struct {
	eventClient eventClient
}

func New(eventClient eventClient) *repository {
	return &repository{
		eventClient: eventClient,
	}
}

func (r *repository) CreateEventRule(ctx context.Context, auctionID string, startDate time.Time) error {
	year, month, day := startDate.Date()
	hour, min, _ := startDate.Clock()
	_, err := r.eventClient.PutRule(ctx, &cloudwatchevents.PutRuleInput{
		Name:               aws.String(fmt.Sprintf("auction-event-%s", auctionID)),
		ScheduleExpression: aws.String(fmt.Sprintf("cron(%d %d %d %d ? %d)", min, hour, day, month, year)),
	})
	if err != nil {
		return err
	}

	eventInput, err := json.Marshal(models.AuctionEvent{
		AuctionID: auctionID,
		Stage:     "STAGE_ACCEPTING_BIDS",
		EndDate:   startDate,
	})
	if err != nil {
		return err
	}
	_, err = r.eventClient.PutTargets(ctx, &cloudwatchevents.PutTargetsInput{
		Rule: aws.String(fmt.Sprintf("auction-event-%s", auctionID)),
		Targets: []cloudwatchTypes.Target{
			{
				Arn:   aws.String("arn:aws:lambda:us-east-1:160902899897:function:auctioneer-infra-backend-HandleAuctionFunction-IU4UT7Wy9oKi"),
				Id:    aws.String("auctioneer-infra-backend-HandleAuctionFunction-IU4UT7Wy9oKi"),
				Input: aws.String(string(eventInput)),
			},
		},
	})
	return err
}

func (r *repository) UpdateEventRule(ctx context.Context, auctionID string, newDate time.Time) error {
	year, month, day := newDate.Date()
	hour, min, _ := newDate.Clock()
	_, err := r.eventClient.PutRule(ctx, &cloudwatchevents.PutRuleInput{
		Name:               aws.String(fmt.Sprintf("auction-event-%s", auctionID)),
		ScheduleExpression: aws.String(fmt.Sprintf("cron(%d %d %d %d ? %d)", min, hour, day, month, year)),
	})
	if err != nil {
		return err
	}

	eventInput, err := json.Marshal(models.AuctionEvent{
		AuctionID: auctionID,
		Stage:     "STAGE_AUCTION_ONGOING",
		EndDate:   newDate,
	})
	if err != nil {
		return err
	}
	_, err = r.eventClient.PutTargets(ctx, &cloudwatchevents.PutTargetsInput{
		Rule: aws.String(fmt.Sprintf("auction-event-%s", auctionID)),
		Targets: []cloudwatchTypes.Target{
			{
				Arn:   aws.String("arn:aws:lambda:us-east-1:160902899897:function:auctioneer-infra-backend-HandleAuctionFunction-IU4UT7Wy9oKi"),
				Id:    aws.String("auctioneer-infra-backend-HandleAuctionFunction-IU4UT7Wy9oKi"),
				Input: aws.String(string(eventInput)),
			},
		},
	})
	return err
}

func (r *repository) DeleteEventRule(ctx context.Context, auctionID string) error {
	_, err := r.eventClient.RemoveTargets(ctx, &cloudwatchevents.RemoveTargetsInput{
		Rule: aws.String(fmt.Sprintf("auction-event-%s", auctionID)),
		Ids:  []string{"auctioneer-infra-backend-HandleAuctionFunction-IU4UT7Wy9oKi"},
	})
	if err != nil {
		return err
	}
	_, err = r.eventClient.DeleteRule(ctx, &cloudwatchevents.DeleteRuleInput{
		Name: aws.String(fmt.Sprintf("auction-event-%s", auctionID)),
	})
	return err
}

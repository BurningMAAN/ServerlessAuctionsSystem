package eventbridge

import (
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
}

type repository struct {
	eventClient eventClient
}

func New(eventClient eventClient) *repository {
	return &repository{
		eventClient: eventClient,
	}
}

type AuctionEvent struct {
	AuctionID string    `json:"id"`
	Stage     string    `json:"stage"`
	EndDate   time.Time `json:"endDate"`
}

func (r *repository) UpdateEventRule(ctx context.Context, auctionID string) error {
	_, err := r.eventClient.PutRule(ctx, &cloudwatchevents.PutRuleInput{
		Name:               aws.String(fmt.Sprintf("auction-event-%s", auctionID)),
		ScheduleExpression: aws.String("rate(5 minutes)"),
	})

	eventInput, err := json.Marshal(AuctionEvent{
		AuctionID: auctionID,
		Stage:     "STAGE_AUCTION_ONGOING",
		EndDate:   time.Now().Add(33 * time.Second),
	})
	if err != nil {
		return err
	}
	_, err = r.eventClient.PutTargets(ctx, &cloudwatchevents.PutTargetsInput{
		Rule: aws.String(fmt.Sprintf("auction-event-%s", auctionID)),
		Targets: []cloudwatchTypes.Target{
			{
				Arn:   aws.String("arn:aws:lambda:us-east-1:102336894219:function:test-backend-HandleAuctionFunction-Oa1T2FivSffq"),
				Id:    aws.String("test-backend-HandleAuctionFunction-Oa1T2FivSffq"),
				Input: aws.String(string(eventInput)),
			},
		},
	})
	return err
}

func (r *repository) DeleteEventRule(ctx context.Context, auctionID string) error {
	_, err := r.eventClient.RemoveTargets(ctx, &cloudwatchevents.RemoveTargetsInput{
		Rule: aws.String(fmt.Sprintf("auction-event-%s", auctionID)),
		Ids:  []string{"test-backend-HandleAuctionFunction-Oa1T2FivSffq"},
	})
	_, err = r.eventClient.DeleteRule(ctx, &cloudwatchevents.DeleteRuleInput{
		Name: aws.String(fmt.Sprintf("auction-event-%s", auctionID)),
	})
	return err
}

package eventbridge

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go/aws"
)

type eventClient interface {
	PutRule(ctx context.Context, params *cloudwatchevents.PutRuleInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.PutRuleOutput, error)
}

type repository struct {
	eventClient eventClient
}

func New(eventClient eventClient) *repository {
	return &repository{
		eventClient: eventClient,
	}
}

func (r *repository) UpdateEventRule(ctx context.Context, auctionID string) error {
	_, err := r.eventClient.PutRule(ctx, &cloudwatchevents.PutRuleInput{
		Name:               aws.String(fmt.Sprintf("auction-event-%s", auctionID)),
		ScheduleExpression: aws.String("rate(5 minutes)"),
	})
	return err
}

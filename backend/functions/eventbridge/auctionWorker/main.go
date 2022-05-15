//go:build !test
// +build !test

package main

import (
	"context"
	"log"

	auctionsRepository "auctionsPlatform/repositories/auction"
	"auctionsPlatform/repositories/eventbridge"

	bidRepo "auctionsPlatform/repositories/bid"
	userRepo "auctionsPlatform/repositories/user"

	"github.com/aws/aws-lambda-go/lambda"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/kelseyhightower/envconfig"
)

type lambdaConfig struct {
	TableName string `envconfig:"DYNAMODB_TABLE" required:"true"`
}

func main() {
	var cfg lambdaConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("failed to read environment variables: %v", err)
	}

	awsCfg, err := awsconfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	db := dynamodb.NewFromConfig(awsCfg)
	auctionRepository := auctionsRepository.New(cfg.TableName, db)
	userRepository := userRepo.New(cfg.TableName, db)
	bidRepository := bidRepo.New(cfg.TableName, db)
	clClient := cloudwatchevents.NewFromConfig(awsCfg)
	h := handler{
		auctionRepo:     auctionRepository,
		eventRepository: eventbridge.New(clClient),
		userRepository:  userRepository,
		bidRepository:   bidRepository,
	}

	lambda.Start(h.HandleAuction)
}

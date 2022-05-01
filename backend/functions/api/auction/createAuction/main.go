//go:build !test
// +build !test

package main

import (
	"context"
	"log"

	auctionsRepository "auctionsPlatform/repositories/auction"
	"auctionsPlatform/repositories/eventbridge"
	itemsRepository "auctionsPlatform/repositories/item"
	auctionsService "auctionsPlatform/services/auction"
	itemService "auctionsPlatform/services/item"

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
	clClient := cloudwatchevents.NewFromConfig(awsCfg)
	auctionRepository := auctionsRepository.New(cfg.TableName, db)
	itemRepository := itemsRepository.New(cfg.TableName, db)

	c := handler{
		auctionService: auctionsService.New(auctionRepository, itemService.New(itemRepository, auctionRepository), eventbridge.New(clClient)),
	}
	lambda.Start(c.CreateAuction)
}

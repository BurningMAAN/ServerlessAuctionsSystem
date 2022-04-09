//go:build !test
// +build !test

package main

import (
	"context"
	"log"

	auctionsRepository "auctionsPlatform/repositories/auction"
	auctionsService "auctionsPlatform/services/auction"

	"github.com/aws/aws-lambda-go/lambda"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/kelseyhightower/envconfig"
)

type lambdaConfig struct {
	LoanServiceTableName string `envconfig:"DYNAMODB_LOAN_DATA_TABLE" required:"true"`
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
	auctionRepository := auctionsRepository.New(db)

	c := handler{
		auctionService: auctionsService.New(auctionRepository),
	}
	lambda.Start(c.CreateAuction)
}

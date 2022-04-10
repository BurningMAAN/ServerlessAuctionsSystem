//go:build !test
// +build !test

package main

import (
	"context"
	"log"

	userRepo "auctionsPlatform/repositories/user"
	authorizerSvc "auctionsPlatform/services/authorizer"

	"github.com/aws/aws-lambda-go/lambda"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
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
	userRepository := userRepo.New(cfg.TableName, db)

	c := handler{
		authorizerService: authorizerSvc.New(userRepository),
	}
	lambda.Start(c.Authorize)
}
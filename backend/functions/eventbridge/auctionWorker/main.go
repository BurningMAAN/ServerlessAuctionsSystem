//go:build !test
// +build !test

package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
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

	// awsCfg, err := awsconfig.LoadDefaultConfig(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	lambda.Start(HandleAuction)
}

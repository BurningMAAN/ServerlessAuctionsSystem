package localdb

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ory/dockertest"
)

const (
	dynamoDbDockerImage = "amazon/dynamodb-local"
	dynamoDbDockerTag   = "latest"
	dynamoDbDockerPort  = "8000"
)

// DB contains DynamoDB client and pointers to created Docker resources
type DB struct {
	TableName string
	DB        *dynamodb.Client

	// pool and container are needed to stop DB instace and clean up
	pool      *dockertest.Pool
	container *dockertest.Resource
}

// TableSchemaProvider is a function which is passed to New to direct how to create DynamoDB table
type TableSchemaProvider func(tableName string) (dynamodb.CreateTableInput, error)

// New runs DynamoDB in Docker container and creates a table using loan-service schema
func New(schemaProvider TableSchemaProvider) (*DB, error) {
	// connect to Docker daemon using default configuration (on macos/linux uses unix socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("connect to docker daemon: %w", err)
	}

	// pulls an image, creates a container based on it and runs it
	container, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository:   dynamoDbDockerImage,
		Tag:          dynamoDbDockerTag,
		ExposedPorts: []string{dynamoDbDockerPort},
	})
	if err != nil {
		return nil, fmt.Errorf("start database image: %v", err)
	}

	// overwrite AWS SDK DynamoDB URL to point to local Docker container
	ddbPort := container.GetPort(fmt.Sprintf("%s/tcp", dynamoDbDockerPort))
	ddbURL := fmt.Sprintf("http://127.0.0.1:%s", ddbPort)
	localResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if service == dynamodb.ServiceID {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           ddbURL,
				SigningRegion: "us-east-1",
			}, nil
		}
		// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	// use fake credentials (DynamoDB in Docker does not care about credentials validity but they are still needed)
	credProvider := aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     "ak",
			SecretAccessKey: "sk",
		}, nil
	})

	// use custom retry configuration (with 5 retries instead of default 3) to ensure we connect if Docker container is slow to start
	retryer := func() aws.Retryer {
		return retry.AddWithMaxAttempts(retry.NewStandard(), 5)
	}

	// create client using our custom Endpoint and Credentials configuration
	ctx := context.Background()
	config, err := config.LoadDefaultConfig(ctx,
		config.WithEndpointResolver(localResolver),
		config.WithCredentialsProvider(credProvider),
		config.WithRetryer(retryer),
	)
	if err != nil {
		log.Fatal(err)
	}
	db := dynamodb.NewFromConfig(config)
	// tableName is hardcoded because each test package will run its own container so no conflicts are expected
	tableName := "test_instance"

	schema, err := schemaProvider(tableName)
	if err != nil {
		return nil, fmt.Errorf("failed to construct table schema: %v", err)
	}

	_, err = db.CreateTable(ctx, &schema)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	localDB := &DB{
		TableName: tableName,
		DB:        db,
		pool:      pool,
		container: container,
	}

	return localDB, nil
}

// MustNew calls New and fails with fatal error if DB instance cannot be created
func MustNew(schemaProvider TableSchemaProvider) *DB {

	testDB, err := New(schemaProvider)
	if err != nil {
		log.Fatalf("Failed to run DynamoDB Test instance: %v", err)
	}
	return testDB
}

// Stop must be called after tests are finished to stop and remove DynamoDB Docker container
func (db *DB) Stop() error {
	if err := db.pool.Purge(db.container); err != nil {
		return err
	}

	return nil
}

// MustStop calls Stop() and fails with fatal error if DynamoDB container cannot be stopped
func (db *DB) MustStop() {
	err := db.Stop()
	if err != nil {
		log.Fatalf("Failed to run DynamoDB Test instance: %v", err)
	}
}

// SimpleTable returns TableSchemaProvider which creates schema with PK and SK only
func SimpleTable() TableSchemaProvider {
	return func(tableName string) (dynamodb.CreateTableInput, error) {
		return dynamodb.CreateTableInput{
			TableName: aws.String("mockTable"),
			AttributeDefinitions: []types.AttributeDefinition{
				{
					AttributeName: aws.String("PK"),
					AttributeType: types.ScalarAttributeTypeS,
				},
				{
					AttributeName: aws.String("SK"),
					AttributeType: types.ScalarAttributeTypeS,
				},
				{
					AttributeName: aws.String("GSI1PK"),
					AttributeType: types.ScalarAttributeTypeS,
				},
				{
					AttributeName: aws.String("GSI1SK"),
					AttributeType: types.ScalarAttributeTypeS,
				},
				{
					AttributeName: aws.String("GSI2PK"),
					AttributeType: types.ScalarAttributeTypeS,
				},
				{
					AttributeName: aws.String("GSI2SK"),
					AttributeType: types.ScalarAttributeTypeS,
				},
			},
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("PK"),
					KeyType:       "HASH",
				},
				{
					AttributeName: aws.String("SK"),
					KeyType:       "RANGE",
				},
			},
			GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
				{
					IndexName: aws.String("GSI1"),
					Projection: &types.Projection{
						ProjectionType: types.ProjectionTypeAll,
					},
					KeySchema: []types.KeySchemaElement{
						{
							AttributeName: aws.String("GSI1PK"),
							KeyType:       "HASH",
						},
						{
							AttributeName: aws.String("GSI1SK"),
							KeyType:       "RANGE",
						},
					},
				},
				{
					IndexName: aws.String("GSI2"),
					Projection: &types.Projection{
						ProjectionType: types.ProjectionTypeAll,
					},
					KeySchema: []types.KeySchemaElement{
						{
							AttributeName: aws.String("GSI2PK"),
							KeyType:       "HASH",
						},
						{
							AttributeName: aws.String("GSI2SK"),
							KeyType:       "RANGE",
						},
					},
				},
			},
			BillingMode: types.BillingModePayPerRequest,
		}, nil
	}
}

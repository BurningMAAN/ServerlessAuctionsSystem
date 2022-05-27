package localdb

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/awslabs/goformation/v4"
	cfdynamodb "github.com/awslabs/goformation/v4/cloudformation/dynamodb"
)

// SchemaFromCloudformation returns CreateTableInput constructed from schema defined in CF template
func SchemaFromCloudformation(cfTemplateLocation, tableResourceName string) TableSchemaProvider {
	return func(tableName string) (dynamodb.CreateTableInput, error) {
		template, err := goformation.Open(cfTemplateLocation)
		if err != nil {
			return dynamodb.CreateTableInput{}, fmt.Errorf("read Cloudformation template %s: %w", cfTemplateLocation, err)
		}

		table, err := template.GetDynamoDBTableWithName(tableResourceName)
		if err != nil {
			return dynamodb.CreateTableInput{}, fmt.Errorf("get DynamoDB table resource %s from template: %v", tableResourceName, err)
		}

		createTableInput := dynamodb.CreateTableInput{
			AttributeDefinitions:   convertAttributeDefinitions(table.AttributeDefinitions),
			KeySchema:              convertKeySchema(table.KeySchema),
			TableName:              aws.String(tableName),
			BillingMode:            types.BillingMode(table.BillingMode),
			GlobalSecondaryIndexes: convertGSIs(table.GlobalSecondaryIndexes),
		}

		return createTableInput, nil
	}
}

func convertAttributeDefinitions(cfDefinitions []cfdynamodb.Table_AttributeDefinition) []types.AttributeDefinition {
	attributeDefs := make([]types.AttributeDefinition, 0, len(cfDefinitions))
	for _, attrDef := range cfDefinitions {
		attributeDefs = append(attributeDefs, types.AttributeDefinition{
			AttributeName: aws.String(attrDef.AttributeName),
			AttributeType: types.ScalarAttributeType(attrDef.AttributeType),
		})
	}
	return attributeDefs
}

func convertKeySchema(cfKeySchema []cfdynamodb.Table_KeySchema) []types.KeySchemaElement {
	keySchema := make([]types.KeySchemaElement, 0, len(cfKeySchema))
	for _, elem := range cfKeySchema {
		keySchema = append(keySchema, types.KeySchemaElement{
			AttributeName: aws.String(elem.AttributeName),
			KeyType:       types.KeyType(elem.KeyType),
		})
	}
	return keySchema
}

func convertGSIs(cfGSIs []cfdynamodb.Table_GlobalSecondaryIndex) []types.GlobalSecondaryIndex {
	gsis := make([]types.GlobalSecondaryIndex, 0, len(cfGSIs))
	for _, gsi := range cfGSIs {
		gsis = append(gsis, types.GlobalSecondaryIndex{
			IndexName: aws.String(gsi.IndexName),
			KeySchema: convertKeySchema(gsi.KeySchema),
			Projection: &types.Projection{
				ProjectionType: types.ProjectionType(gsi.Projection.ProjectionType),
			},
		})
	}
	return gsis
}

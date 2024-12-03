package dynamodb

import (
	"go.khulnasoft.com/pkg/iac/providers/aws/dynamodb"
	"go.khulnasoft.com/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a dynamodb instance
func Adapt(cfFile parser.FileContext) dynamodb.DynamoDB {
	return dynamodb.DynamoDB{
		DAXClusters: getClusters(cfFile),
	}
}

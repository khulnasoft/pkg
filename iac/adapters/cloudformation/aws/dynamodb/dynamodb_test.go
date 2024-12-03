package dynamodb

import (
	"testing"

	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/testutil"
	"go.khulnasoft.com/pkg/iac/providers/aws/dynamodb"
	"go.khulnasoft.com/pkg/iac/types"
)

func TestAdapt(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		expected dynamodb.DynamoDB
	}{
		{
			name: "complete",
			source: `AWSTemplateFormatVersion: '2010-09-09'
Resources:
  daxCluster:
    Type: AWS::DAX::Cluster
    Properties:
      SSESpecification:
        SSEEnabled: true
`,
			expected: dynamodb.DynamoDB{
				DAXClusters: []dynamodb.DAXCluster{
					{
						ServerSideEncryption: dynamodb.ServerSideEncryption{
							Enabled: types.BoolTest(true),
						},
					},
				},
			},
		},
		{
			name: "empty",
			source: `AWSTemplateFormatVersion: 2010-09-09
Resources:
  daxCluster:
    Type: AWS::DAX::Cluster
  `,
			expected: dynamodb.DynamoDB{
				DAXClusters: []dynamodb.DAXCluster{{}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.AdaptAndCompare(t, tt.source, tt.expected, Adapt)
		})
	}
}

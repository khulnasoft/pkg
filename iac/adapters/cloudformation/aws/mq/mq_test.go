package mq

import (
	"testing"

	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/testutil"
	"go.khulnasoft.com/pkg/iac/providers/aws/mq"
	"go.khulnasoft.com/pkg/iac/types"
)

func TestAdapt(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		expected mq.MQ
	}{
		{
			name: "complete",
			source: `AWSTemplateFormatVersion: '2010-09-09'
Resources: 
  BasicBroker:
    Type: "AWS::AmazonMQ::Broker"
    Properties: 
      PubliclyAccessible: true
      Logs:
        Audit: true
        General: true
`,
			expected: mq.MQ{
				Brokers: []mq.Broker{
					{
						PublicAccess: types.BoolTest(true),
						Logging: mq.Logging{
							Audit:   types.BoolTest(true),
							General: types.BoolTest(true),
						},
					},
				},
			},
		},
		{
			name: "empty",
			source: `AWSTemplateFormatVersion: 2010-09-09
Resources: 
  BasicBroker:
    Type: "AWS::AmazonMQ::Broker"
  `,
			expected: mq.MQ{
				Brokers: []mq.Broker{{}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.AdaptAndCompare(t, tt.source, tt.expected, Adapt)
		})
	}
}

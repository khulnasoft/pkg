package kinesis

import (
	"testing"

	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/testutil"
	"go.khulnasoft.com/pkg/iac/providers/aws/kinesis"
	"go.khulnasoft.com/pkg/iac/types"
)

func TestAdapt(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		expected kinesis.Kinesis
	}{
		{
			name: "complete",
			source: `AWSTemplateFormatVersion: '2010-09-09'
Resources:
  MyStream:
    Type: 'AWS::Kinesis::Stream'
    Properties:
      StreamEncryption:
        EncryptionType: KMS
        KeyId: key
`,
			expected: kinesis.Kinesis{
				Streams: []kinesis.Stream{
					{
						Encryption: kinesis.Encryption{
							Type:     types.StringTest("KMS"),
							KMSKeyID: types.StringTest("key"),
						},
					},
				},
			},
		},
		{
			name: "empty",
			source: `AWSTemplateFormatVersion: 2010-09-09
Resources:
  MyStream:
    Type: 'AWS::Kinesis::Stream'
  `,
			expected: kinesis.Kinesis{
				Streams: []kinesis.Stream{{}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.AdaptAndCompare(t, tt.source, tt.expected, Adapt)
		})
	}
}

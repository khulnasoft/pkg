package sns

import (
	"go.khulnasoft.com/pkg/iac/providers/aws/sns"
	"go.khulnasoft.com/pkg/iac/scanners/cloudformation/parser"
	"go.khulnasoft.com/pkg/iac/types"
)

func getTopics(ctx parser.FileContext) (topics []sns.Topic) {
	for _, r := range ctx.GetResourcesByType("AWS::SNS::Topic") {

		topic := sns.Topic{
			Metadata: r.Metadata(),
			ARN:      types.StringDefault("", r.Metadata()),
			Encryption: sns.Encryption{
				Metadata: r.Metadata(),
				KMSKeyID: r.GetStringProperty("KmsMasterKeyId"),
			},
		}

		topics = append(topics, topic)
	}
	return topics
}

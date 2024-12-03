package sqs

import (
	"go.khulnasoft.com/pkg/iac/providers/aws/sqs"
	"go.khulnasoft.com/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an SQS instance
func Adapt(cfFile parser.FileContext) sqs.SQS {
	return sqs.SQS{
		Queues: getQueues(cfFile),
	}
}

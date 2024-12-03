package cloudwatch

import (
	"go.khulnasoft.com/pkg/iac/providers/aws/cloudwatch"
	"go.khulnasoft.com/pkg/iac/scanners/cloudformation/parser"
)

func getLogGroups(ctx parser.FileContext) (logGroups []cloudwatch.LogGroup) {

	logGroupResources := ctx.GetResourcesByType("AWS::Logs::LogGroup")

	for _, r := range logGroupResources {
		group := cloudwatch.LogGroup{
			Metadata:        r.Metadata(),
			Name:            r.GetStringProperty("LogGroupName"),
			KMSKeyID:        r.GetStringProperty("KmsKeyId"),
			RetentionInDays: r.GetIntProperty("RetentionInDays"),
		}
		logGroups = append(logGroups, group)
	}

	return logGroups
}

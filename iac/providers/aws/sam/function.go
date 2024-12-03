package sam

import (
	"go.khulnasoft.com/pkg/iac/providers/aws/iam"
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

type Function struct {
	Metadata        iacTypes.Metadata
	FunctionName    iacTypes.StringValue
	Tracing         iacTypes.StringValue
	ManagedPolicies []iacTypes.StringValue
	Policies        []iam.Policy
}

const (
	TracingModePassThrough = "PassThrough"
	TracingModeActive      = "Active"
)

type Permission struct {
	Metadata  iacTypes.Metadata
	Principal iacTypes.StringValue
	SourceARN iacTypes.StringValue
}

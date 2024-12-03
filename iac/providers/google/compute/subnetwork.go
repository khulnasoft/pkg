package compute

import (
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

type SubNetwork struct {
	Metadata       iacTypes.Metadata
	Name           iacTypes.StringValue
	Purpose        iacTypes.StringValue
	EnableFlowLogs iacTypes.BoolValue
}

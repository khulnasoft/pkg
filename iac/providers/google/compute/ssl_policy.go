package compute

import (
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

type SSLPolicy struct {
	Metadata          iacTypes.Metadata
	Name              iacTypes.StringValue
	Profile           iacTypes.StringValue
	MinimumTLSVersion iacTypes.StringValue
}

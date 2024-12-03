package network

import (
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

type VpnGateway struct {
	Metadata      iacTypes.Metadata
	SecurityGroup iacTypes.StringValue
}

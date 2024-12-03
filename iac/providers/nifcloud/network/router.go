package network

import (
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

type Router struct {
	Metadata          iacTypes.Metadata
	SecurityGroup     iacTypes.StringValue
	NetworkInterfaces []NetworkInterface
}

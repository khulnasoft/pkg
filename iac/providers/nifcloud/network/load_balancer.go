package network

import (
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

type LoadBalancer struct {
	Metadata  iacTypes.Metadata
	Listeners []LoadBalancerListener
}

type LoadBalancerListener struct {
	Metadata  iacTypes.Metadata
	Protocol  iacTypes.StringValue
	TLSPolicy iacTypes.StringValue
}

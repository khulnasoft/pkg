package network

import (
	"go.khulnasoft.com/pkg/iac/providers/nifcloud/network"
	"go.khulnasoft.com/pkg/iac/terraform"
)

func adaptVpnGateways(modules terraform.Modules) []network.VpnGateway {
	var vpnGateways []network.VpnGateway

	for _, resource := range modules.GetResourcesByType("nifcloud_vpn_gateway") {
		vpnGateways = append(vpnGateways, adaptVpnGateway(resource))
	}
	return vpnGateways
}

func adaptVpnGateway(resource *terraform.Block) network.VpnGateway {
	return network.VpnGateway{
		Metadata:      resource.GetMetadata(),
		SecurityGroup: resource.GetAttribute("security_group").AsStringValueOrDefault("", resource),
	}
}

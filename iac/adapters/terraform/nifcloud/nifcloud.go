package nifcloud

import (
	"go.khulnasoft.com/pkg/iac/adapters/terraform/nifcloud/computing"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/nifcloud/dns"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/nifcloud/nas"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/nifcloud/network"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/nifcloud/rdb"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/nifcloud/sslcertificate"
	"go.khulnasoft.com/pkg/iac/providers/nifcloud"
	"go.khulnasoft.com/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) nifcloud.Nifcloud {
	return nifcloud.Nifcloud{
		Computing:      computing.Adapt(modules),
		DNS:            dns.Adapt(modules),
		NAS:            nas.Adapt(modules),
		Network:        network.Adapt(modules),
		RDB:            rdb.Adapt(modules),
		SSLCertificate: sslcertificate.Adapt(modules),
	}
}

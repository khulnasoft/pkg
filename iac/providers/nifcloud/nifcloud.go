package nifcloud

import (
	"go.khulnasoft.com/pkg/iac/providers/nifcloud/computing"
	"go.khulnasoft.com/pkg/iac/providers/nifcloud/dns"
	"go.khulnasoft.com/pkg/iac/providers/nifcloud/nas"
	"go.khulnasoft.com/pkg/iac/providers/nifcloud/network"
	"go.khulnasoft.com/pkg/iac/providers/nifcloud/rdb"
	"go.khulnasoft.com/pkg/iac/providers/nifcloud/sslcertificate"
)

type Nifcloud struct {
	Computing      computing.Computing
	DNS            dns.DNS
	NAS            nas.NAS
	Network        network.Network
	RDB            rdb.RDB
	SSLCertificate sslcertificate.SSLCertificate
}

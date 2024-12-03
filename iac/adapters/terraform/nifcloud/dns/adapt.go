package dns

import (
	"go.khulnasoft.com/pkg/iac/providers/nifcloud/dns"
	"go.khulnasoft.com/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) dns.DNS {
	return dns.DNS{
		Records: adaptRecords(modules),
	}
}

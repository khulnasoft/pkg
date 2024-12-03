package sslcertificate

import (
	"go.khulnasoft.com/pkg/iac/providers/nifcloud/sslcertificate"
	"go.khulnasoft.com/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) sslcertificate.SSLCertificate {
	return sslcertificate.SSLCertificate{
		ServerCertificates: adaptServerCertificates(modules),
	}
}

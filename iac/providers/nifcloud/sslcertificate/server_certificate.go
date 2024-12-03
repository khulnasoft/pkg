package sslcertificate

import (
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

type ServerCertificate struct {
	Metadata   iacTypes.Metadata
	Expiration iacTypes.TimeValue
}

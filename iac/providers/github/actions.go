package github

import (
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

type Action struct {
	Metadata           iacTypes.Metadata
	EnvironmentSecrets []EnvironmentSecret
}

type EnvironmentSecret struct {
	Metadata       iacTypes.Metadata
	Repository     iacTypes.StringValue
	Environment    iacTypes.StringValue
	SecretName     iacTypes.StringValue
	PlainTextValue iacTypes.StringValue
	EncryptedValue iacTypes.StringValue
}

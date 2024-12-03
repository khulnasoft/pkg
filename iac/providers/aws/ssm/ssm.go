package ssm

import (
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

type SSM struct {
	Secrets []Secret
}

type Secret struct {
	Metadata iacTypes.Metadata
	KMSKeyID iacTypes.StringValue
}

const DefaultKMSKeyID = "alias/aws/secretsmanager"

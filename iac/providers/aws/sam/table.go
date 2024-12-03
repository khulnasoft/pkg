package sam

import (
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

type SimpleTable struct {
	Metadata         iacTypes.Metadata
	TableName        iacTypes.StringValue
	SSESpecification SSESpecification
}

type SSESpecification struct {
	Metadata iacTypes.Metadata

	Enabled        iacTypes.BoolValue
	KMSMasterKeyID iacTypes.StringValue
}

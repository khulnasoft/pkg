package kinesis

import (
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

type Kinesis struct {
	Streams []Stream
}

type Stream struct {
	Metadata   iacTypes.Metadata
	Encryption Encryption
}

const (
	EncryptionTypeKMS = "KMS"
)

type Encryption struct {
	Metadata iacTypes.Metadata
	Type     iacTypes.StringValue
	KMSKeyID iacTypes.StringValue
}

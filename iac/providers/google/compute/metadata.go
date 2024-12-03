package compute

import (
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

type ProjectMetadata struct {
	Metadata      iacTypes.Metadata
	EnableOSLogin iacTypes.BoolValue
}

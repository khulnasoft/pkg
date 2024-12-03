package compute

import (
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

type Compute struct {
	Instances []Instance
}

type Instance struct {
	Metadata iacTypes.Metadata
	UserData iacTypes.StringValue // not b64 encoded pls
}

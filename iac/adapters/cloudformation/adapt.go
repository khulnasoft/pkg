package cloudformation

import (
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws"
	"go.khulnasoft.com/pkg/iac/scanners/cloudformation/parser"
	"go.khulnasoft.com/pkg/iac/state"
)

// Adapt adapts the Cloudformation instance
func Adapt(cfFile parser.FileContext) *state.State {
	return &state.State{
		AWS: aws.Adapt(cfFile),
	}
}

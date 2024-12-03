package dockerfile

import (
	"go.khulnasoft.com/pkg/iac/scanners/dockerfile/parser"
	"go.khulnasoft.com/pkg/iac/scanners/generic"
	"go.khulnasoft.com/pkg/iac/scanners/options"
	"go.khulnasoft.com/pkg/iac/types"
)

func NewScanner(opts ...options.ScannerOption) *generic.GenericScanner {
	return generic.NewScanner("Dockerfile", types.SourceDockerfile, generic.ParseFunc(parser.Parse), opts...)
}

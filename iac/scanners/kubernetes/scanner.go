package kubernetes

import (
	"context"
	"io"

	"go.khulnasoft.com/pkg/iac/scanners/generic"
	"go.khulnasoft.com/pkg/iac/scanners/kubernetes/parser"
	"go.khulnasoft.com/pkg/iac/scanners/options"
	"go.khulnasoft.com/pkg/iac/types"
)

func NewScanner(opts ...options.ScannerOption) *generic.GenericScanner {
	return generic.NewScanner("Kubernetes", types.SourceKubernetes, generic.ParseFunc(parse), opts...)
}

func parse(ctx context.Context, r io.Reader, path string) (any, error) {
	return parser.Parse(ctx, r, path)
}

package lambda

import (
	"github.com/aquasecurity/defsec/pkg/providers/aws/lambda"
	"github.com/khulnasoft/tunnel-iac/pkg/scanners/cloudformation/parser"
)

// Adapt ...
func Adapt(cfFile parser.FileContext) lambda.Lambda {
	return lambda.Lambda{
		Functions: getFunctions(cfFile),
	}
}

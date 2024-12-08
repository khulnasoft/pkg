package kinesis

import (
	"github.com/aquasecurity/defsec/pkg/providers/aws/kinesis"
	"github.com/khulnasoft/tunnel-iac/pkg/scanners/cloudformation/parser"
)

// Adapt ...
func Adapt(cfFile parser.FileContext) kinesis.Kinesis {
	return kinesis.Kinesis{
		Streams: getStreams(cfFile),
	}
}

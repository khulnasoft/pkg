package cloudfront

import (
	"go.khulnasoft.com/pkg/iac/providers/aws/cloudfront"
	"go.khulnasoft.com/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a CloudFront instance
func Adapt(cfFile parser.FileContext) cloudfront.Cloudfront {
	return cloudfront.Cloudfront{
		Distributions: getDistributions(cfFile),
	}
}

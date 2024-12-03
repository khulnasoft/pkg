package elasticsearch

import (
	"go.khulnasoft.com/pkg/iac/providers/aws/elasticsearch"
	"go.khulnasoft.com/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an ElasticSearch instance
func Adapt(cfFile parser.FileContext) elasticsearch.Elasticsearch {
	return elasticsearch.Elasticsearch{
		Domains: getDomains(cfFile),
	}
}

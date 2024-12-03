package documentdb

import (
	"go.khulnasoft.com/pkg/iac/providers/aws/documentdb"
	"go.khulnasoft.com/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adaps a documentDB instance
func Adapt(cfFile parser.FileContext) documentdb.DocumentDB {
	return documentdb.DocumentDB{
		Clusters: getClusters(cfFile),
	}
}

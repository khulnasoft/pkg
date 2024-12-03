package rds

import (
	"go.khulnasoft.com/pkg/iac/providers/aws/rds"
	"go.khulnasoft.com/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an RDS instance
func Adapt(cfFile parser.FileContext) rds.RDS {
	clusters, orphans := getClustersAndInstances(cfFile)
	return rds.RDS{
		Instances:       orphans,
		Clusters:        clusters,
		Classic:         getClassic(cfFile),
		ParameterGroups: getParameterGroups(cfFile),
		Snapshots:       nil,
	}
}

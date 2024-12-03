package sam

import (
	"go.khulnasoft.com/pkg/iac/providers/aws/sam"
	"go.khulnasoft.com/pkg/iac/scanners/cloudformation/parser"
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

func getSimpleTables(cfFile parser.FileContext) (tables []sam.SimpleTable) {

	tableResources := cfFile.GetResourcesByType("AWS::Serverless::SimpleTable")
	for _, r := range tableResources {
		table := sam.SimpleTable{
			Metadata:         r.Metadata(),
			TableName:        r.GetStringProperty("TableName"),
			SSESpecification: getSSESpecification(r),
		}

		tables = append(tables, table)
	}

	return tables
}

func getSSESpecification(r *parser.Resource) sam.SSESpecification {
	if sse := r.GetProperty("SSESpecification"); sse.IsNotNil() {
		return sam.SSESpecification{
			Metadata:       sse.Metadata(),
			Enabled:        sse.GetBoolProperty("SSEEnabled"),
			KMSMasterKeyID: sse.GetStringProperty("KMSMasterKeyId"),
		}
	}

	return sam.SSESpecification{
		Metadata:       r.Metadata(),
		Enabled:        iacTypes.BoolDefault(false, r.Metadata()),
		KMSMasterKeyID: iacTypes.StringDefault("", r.Metadata()),
	}
}

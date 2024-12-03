package google

import (
	"go.khulnasoft.com/pkg/iac/adapters/terraform/google/bigquery"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/google/compute"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/google/dns"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/google/gke"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/google/iam"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/google/kms"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/google/sql"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/google/storage"
	"go.khulnasoft.com/pkg/iac/providers/google"
	"go.khulnasoft.com/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) google.Google {
	return google.Google{
		BigQuery: bigquery.Adapt(modules),
		Compute:  compute.Adapt(modules),
		DNS:      dns.Adapt(modules),
		GKE:      gke.Adapt(modules),
		KMS:      kms.Adapt(modules),
		IAM:      iam.Adapt(modules),
		SQL:      sql.Adapt(modules),
		Storage:  storage.Adapt(modules),
	}
}

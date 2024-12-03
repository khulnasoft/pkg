package google

import (
	"go.khulnasoft.com/pkg/iac/providers/google/bigquery"
	"go.khulnasoft.com/pkg/iac/providers/google/compute"
	"go.khulnasoft.com/pkg/iac/providers/google/dns"
	"go.khulnasoft.com/pkg/iac/providers/google/gke"
	"go.khulnasoft.com/pkg/iac/providers/google/iam"
	"go.khulnasoft.com/pkg/iac/providers/google/kms"
	"go.khulnasoft.com/pkg/iac/providers/google/sql"
	"go.khulnasoft.com/pkg/iac/providers/google/storage"
)

type Google struct {
	BigQuery bigquery.BigQuery
	Compute  compute.Compute
	DNS      dns.DNS
	GKE      gke.GKE
	KMS      kms.KMS
	IAM      iam.IAM
	SQL      sql.SQL
	Storage  storage.Storage
}

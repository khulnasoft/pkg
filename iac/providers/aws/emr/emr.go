package emr

import (
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

type EMR struct {
	Clusters              []Cluster
	SecurityConfiguration []SecurityConfiguration
}

type Cluster struct {
	Metadata iacTypes.Metadata
	Settings ClusterSettings
}

type ClusterSettings struct {
	Metadata     iacTypes.Metadata
	Name         iacTypes.StringValue
	ReleaseLabel iacTypes.StringValue
	ServiceRole  iacTypes.StringValue
}

type SecurityConfiguration struct {
	Metadata      iacTypes.Metadata
	Name          iacTypes.StringValue
	Configuration iacTypes.StringValue
}

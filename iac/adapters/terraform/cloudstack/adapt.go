package cloudstack

import (
	"go.khulnasoft.com/pkg/iac/adapters/terraform/cloudstack/compute"
	"go.khulnasoft.com/pkg/iac/providers/cloudstack"
	"go.khulnasoft.com/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) cloudstack.CloudStack {
	return cloudstack.CloudStack{
		Compute: compute.Adapt(modules),
	}
}

package terraform

import (
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/azure"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/cloudstack"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/digitalocean"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/github"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/google"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/kubernetes"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/nifcloud"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/openstack"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/oracle"
	"go.khulnasoft.com/pkg/iac/state"
	"go.khulnasoft.com/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) *state.State {
	return &state.State{
		AWS:          aws.Adapt(modules),
		Azure:        azure.Adapt(modules),
		CloudStack:   cloudstack.Adapt(modules),
		DigitalOcean: digitalocean.Adapt(modules),
		GitHub:       github.Adapt(modules),
		Google:       google.Adapt(modules),
		Kubernetes:   kubernetes.Adapt(modules),
		Nifcloud:     nifcloud.Adapt(modules),
		OpenStack:    openstack.Adapt(modules),
		Oracle:       oracle.Adapt(modules),
	}
}

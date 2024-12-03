package digitalocean

import (
	"go.khulnasoft.com/pkg/iac/adapters/terraform/digitalocean/compute"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/digitalocean/spaces"
	"go.khulnasoft.com/pkg/iac/providers/digitalocean"
	"go.khulnasoft.com/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) digitalocean.DigitalOcean {
	return digitalocean.DigitalOcean{
		Compute: compute.Adapt(modules),
		Spaces:  spaces.Adapt(modules),
	}
}

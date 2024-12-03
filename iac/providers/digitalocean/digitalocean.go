package digitalocean

import (
	"go.khulnasoft.com/pkg/iac/providers/digitalocean/compute"
	"go.khulnasoft.com/pkg/iac/providers/digitalocean/spaces"
)

type DigitalOcean struct {
	Compute compute.Compute
	Spaces  spaces.Spaces
}

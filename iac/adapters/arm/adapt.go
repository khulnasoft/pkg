package arm

import (
	"context"

	"go.khulnasoft.com/pkg/iac/adapters/arm/appservice"
	"go.khulnasoft.com/pkg/iac/adapters/arm/authorization"
	"go.khulnasoft.com/pkg/iac/adapters/arm/compute"
	"go.khulnasoft.com/pkg/iac/adapters/arm/container"
	"go.khulnasoft.com/pkg/iac/adapters/arm/database"
	"go.khulnasoft.com/pkg/iac/adapters/arm/datafactory"
	"go.khulnasoft.com/pkg/iac/adapters/arm/datalake"
	"go.khulnasoft.com/pkg/iac/adapters/arm/keyvault"
	"go.khulnasoft.com/pkg/iac/adapters/arm/monitor"
	"go.khulnasoft.com/pkg/iac/adapters/arm/network"
	"go.khulnasoft.com/pkg/iac/adapters/arm/securitycenter"
	"go.khulnasoft.com/pkg/iac/adapters/arm/storage"
	"go.khulnasoft.com/pkg/iac/adapters/arm/synapse"
	"go.khulnasoft.com/pkg/iac/providers/azure"
	scanner "go.khulnasoft.com/pkg/iac/scanners/azure"
	"go.khulnasoft.com/pkg/iac/state"
)

// Adapt adapts an azure arm instance
func Adapt(ctx context.Context, deployment scanner.Deployment) *state.State {
	return &state.State{
		Azure: adaptAzure(deployment),
	}
}

func adaptAzure(deployment scanner.Deployment) azure.Azure {

	return azure.Azure{
		AppService:     appservice.Adapt(deployment),
		Authorization:  authorization.Adapt(deployment),
		Compute:        compute.Adapt(deployment),
		Container:      container.Adapt(deployment),
		Database:       database.Adapt(deployment),
		DataFactory:    datafactory.Adapt(deployment),
		DataLake:       datalake.Adapt(deployment),
		KeyVault:       keyvault.Adapt(deployment),
		Monitor:        monitor.Adapt(deployment),
		Network:        network.Adapt(deployment),
		SecurityCenter: securitycenter.Adapt(deployment),
		Storage:        storage.Adapt(deployment),
		Synapse:        synapse.Adapt(deployment),
	}

}

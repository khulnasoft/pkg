package azure

import (
	"go.khulnasoft.com/pkg/iac/adapters/terraform/azure/appservice"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/azure/authorization"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/azure/compute"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/azure/container"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/azure/database"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/azure/datafactory"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/azure/datalake"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/azure/keyvault"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/azure/monitor"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/azure/network"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/azure/securitycenter"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/azure/storage"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/azure/synapse"
	"go.khulnasoft.com/pkg/iac/providers/azure"
	"go.khulnasoft.com/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) azure.Azure {
	return azure.Azure{
		AppService:     appservice.Adapt(modules),
		Authorization:  authorization.Adapt(modules),
		Compute:        compute.Adapt(modules),
		Container:      container.Adapt(modules),
		Database:       database.Adapt(modules),
		DataFactory:    datafactory.Adapt(modules),
		DataLake:       datalake.Adapt(modules),
		KeyVault:       keyvault.Adapt(modules),
		Monitor:        monitor.Adapt(modules),
		Network:        network.Adapt(modules),
		SecurityCenter: securitycenter.Adapt(modules),
		Storage:        storage.Adapt(modules),
		Synapse:        synapse.Adapt(modules),
	}
}

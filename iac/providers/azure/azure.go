package azure

import (
	"go.khulnasoft.com/pkg/iac/providers/azure/appservice"
	"go.khulnasoft.com/pkg/iac/providers/azure/authorization"
	"go.khulnasoft.com/pkg/iac/providers/azure/compute"
	"go.khulnasoft.com/pkg/iac/providers/azure/container"
	"go.khulnasoft.com/pkg/iac/providers/azure/database"
	"go.khulnasoft.com/pkg/iac/providers/azure/datafactory"
	"go.khulnasoft.com/pkg/iac/providers/azure/datalake"
	"go.khulnasoft.com/pkg/iac/providers/azure/keyvault"
	"go.khulnasoft.com/pkg/iac/providers/azure/monitor"
	"go.khulnasoft.com/pkg/iac/providers/azure/network"
	"go.khulnasoft.com/pkg/iac/providers/azure/securitycenter"
	"go.khulnasoft.com/pkg/iac/providers/azure/storage"
	"go.khulnasoft.com/pkg/iac/providers/azure/synapse"
)

type Azure struct {
	AppService     appservice.AppService
	Authorization  authorization.Authorization
	Compute        compute.Compute
	Container      container.Container
	Database       database.Database
	DataFactory    datafactory.DataFactory
	DataLake       datalake.DataLake
	KeyVault       keyvault.KeyVault
	Monitor        monitor.Monitor
	Network        network.Network
	SecurityCenter securitycenter.SecurityCenter
	Storage        storage.Storage
	Synapse        synapse.Synapse
}

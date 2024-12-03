package container

import (
	"go.khulnasoft.com/pkg/iac/providers/azure/container"
	"go.khulnasoft.com/pkg/iac/scanners/azure"
)

func Adapt(deployment azure.Deployment) container.Container {
	return container.Container{
		KubernetesClusters: adaptKubernetesClusters(deployment),
	}
}

func adaptKubernetesClusters(deployment azure.Deployment) []container.KubernetesCluster {

	return nil
}

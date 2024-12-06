//go:build wireinject
// +build wireinject

package k8s

import (
	"github.com/google/wire"

	"go.khulnasoft.com/pkg/cache"
)

func initializeScanK8s(localArtifactCache cache.LocalArtifactCache) *ScanKubernetes {
	wire.Build(ScanSuperSet)
	return &ScanKubernetes{}
}

package k8s

import (
	"context"

	"github.com/google/wire"

	ftypes "go.khulnasoft.com/pkg/fanal/types"
	"go.khulnasoft.com/pkg/scanner"
	"go.khulnasoft.com/pkg/scanner/local"
	"go.khulnasoft.com/pkg/types"
)

// ScanSuperSet binds the dependencies for k8s
var ScanSuperSet = wire.NewSet(
	local.SuperSet,
	wire.Bind(new(scanner.Driver), new(local.Scanner)),
	NewScanKubernetes,
)

// ScanKubernetes implements the scanner
type ScanKubernetes struct {
	localScanner local.Scanner
}

// NewScanKubernetes is the factory method for scanner
func NewScanKubernetes(s local.Scanner) *ScanKubernetes {
	return &ScanKubernetes{localScanner: s}
}

// NewKubernetesScanner is the factory method for scanner
func NewKubernetesScanner() *ScanKubernetes {
	return initializeScanK8s(nil)
}

// Scan scans k8s core components and return it findings
func (sk ScanKubernetes) Scan(ctx context.Context, target types.ScanTarget, options types.ScanOptions) (types.Results, ftypes.OS, error) {
	return sk.localScanner.ScanTarget(ctx, target, options)
}

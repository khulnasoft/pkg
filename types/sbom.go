package types

import (
	ftypes "go.khulnasoft.com/pkg/fanal/types"
	"go.khulnasoft.com/pkg/sbom/core"
)

type SBOMSource = string

type SBOM struct {
	Metadata Metadata

	Packages     []ftypes.PackageInfo
	Applications []ftypes.Application

	BOM *core.BOM
}

const (
	SBOMSourceOCI   = SBOMSource("oci")
	SBOMSourceRekor = SBOMSource("rekor")
)

var (
	SBOMSources = []string{
		SBOMSourceOCI,
		SBOMSourceRekor,
	}
)

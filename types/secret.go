package types

import (
	ftypes "go.khulnasoft.com/pkg/fanal/types"
)

type DetectedSecret ftypes.SecretFinding

func (DetectedSecret) findingType() FindingType { return FindingTypeSecret }

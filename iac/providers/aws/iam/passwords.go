package iam

import (
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

type PasswordPolicy struct {
	Metadata             iacTypes.Metadata
	ReusePreventionCount iacTypes.IntValue
	RequireLowercase     iacTypes.BoolValue
	RequireUppercase     iacTypes.BoolValue
	RequireNumbers       iacTypes.BoolValue
	RequireSymbols       iacTypes.BoolValue
	MaxAgeDays           iacTypes.IntValue
	MinimumLength        iacTypes.IntValue
}

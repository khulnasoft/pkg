package datalake

import (
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

type DataLake struct {
	Stores []Store
}

type Store struct {
	Metadata         iacTypes.Metadata
	EnableEncryption iacTypes.BoolValue
}

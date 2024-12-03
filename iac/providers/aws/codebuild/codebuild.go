package codebuild

import (
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

type CodeBuild struct {
	Projects []Project
}

type Project struct {
	Metadata                  iacTypes.Metadata
	ArtifactSettings          ArtifactSettings
	SecondaryArtifactSettings []ArtifactSettings
}

type ArtifactSettings struct {
	Metadata          iacTypes.Metadata
	EncryptionEnabled iacTypes.BoolValue
}

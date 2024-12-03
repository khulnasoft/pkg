package codebuild

import (
	"testing"

	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/testutil"
	"go.khulnasoft.com/pkg/iac/providers/aws/codebuild"
	"go.khulnasoft.com/pkg/iac/types"
)

func TestAdapt(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		expected codebuild.CodeBuild
	}{
		{
			name: "complete",
			source: `AWSTemplateFormatVersion: 2010-09-09
Resources:
  Project:
    Type: AWS::CodeBuild::Project
    Properties:
      Artifacts:
        EncryptionDisabled: true
      SecondaryArtifacts:
        - EncryptionDisabled: true
`,
			expected: codebuild.CodeBuild{
				Projects: []codebuild.Project{
					{
						ArtifactSettings: codebuild.ArtifactSettings{
							EncryptionEnabled: types.BoolTest(false),
						},
						SecondaryArtifactSettings: []codebuild.ArtifactSettings{
							{
								EncryptionEnabled: types.BoolTest(false),
							},
						},
					},
				},
			},
		},
		{
			name: "empty",
			source: `AWSTemplateFormatVersion: 2010-09-09
Resources:
  Project:
    Type: AWS::CodeBuild::Project
  `,
			expected: codebuild.CodeBuild{
				Projects: []codebuild.Project{
					{
						ArtifactSettings: codebuild.ArtifactSettings{
							EncryptionEnabled: types.BoolTest(true),
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.AdaptAndCompare(t, tt.source, tt.expected, Adapt)
		})
	}
}

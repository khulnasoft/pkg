package workspaces

import (
	"testing"

	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/testutil"
	"go.khulnasoft.com/pkg/iac/providers/aws/workspaces"
	"go.khulnasoft.com/pkg/iac/types"
)

func TestAdapt(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		expected workspaces.WorkSpaces
	}{
		{
			name: "complete",
			source: `AWSTemplateFormatVersion: '2010-09-09'
Resources:
  MyWorkSpace:
    Type: AWS::WorkSpaces::Workspace
    Properties:
      RootVolumeEncryptionEnabled: true
      UserVolumeEncryptionEnabled: true
`,
			expected: workspaces.WorkSpaces{
				WorkSpaces: []workspaces.WorkSpace{
					{
						RootVolume: workspaces.Volume{
							Encryption: workspaces.Encryption{
								Enabled: types.BoolTest(true),
							},
						},
						UserVolume: workspaces.Volume{
							Encryption: workspaces.Encryption{
								Enabled: types.BoolTest(true),
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
  MyWorkSpace:
    Type: AWS::WorkSpaces::Workspace
  `,
			expected: workspaces.WorkSpaces{
				WorkSpaces: []workspaces.WorkSpace{{}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.AdaptAndCompare(t, tt.source, tt.expected, Adapt)
		})
	}
}

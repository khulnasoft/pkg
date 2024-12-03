package compute

import (
	"testing"

	"go.khulnasoft.com/tunnel/internal/testutil"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/tftestutil"
	"go.khulnasoft.com/pkg/iac/providers/google/compute"
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

func Test_adaptProjectMetadata(t *testing.T) {
	tests := []struct {
		name      string
		terraform string
		expected  compute.ProjectMetadata
	}{
		{
			name: "defined",
			terraform: `
			resource "google_compute_project_metadata" "example" {
				metadata = {
				  enable-oslogin = true
				}
			  }
`,
			expected: compute.ProjectMetadata{
				Metadata:      iacTypes.NewTestMetadata(),
				EnableOSLogin: iacTypes.Bool(true, iacTypes.NewTestMetadata()),
			},
		},
		{
			name: "defaults",
			terraform: `
			resource "google_compute_project_metadata" "example" {
				metadata = {
				}
			  }
`,
			expected: compute.ProjectMetadata{
				Metadata:      iacTypes.NewTestMetadata(),
				EnableOSLogin: iacTypes.Bool(false, iacTypes.NewTestMetadata()),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			modules := tftestutil.CreateModulesFromSource(t, test.terraform, ".tf")
			adapted := adaptProjectMetadata(modules)
			testutil.AssertDefsecEqual(t, test.expected, adapted)
		})
	}
}

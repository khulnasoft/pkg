package iam

import (
	"testing"

	"go.khulnasoft.com/tunnel/internal/testutil"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/tftestutil"
	"go.khulnasoft.com/pkg/iac/providers/google/iam"
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

func Test_AdaptBinding(t *testing.T) {
	tests := []struct {
		name      string
		terraform string
		expected  iam.Binding
	}{
		{
			name: "defined",
			terraform: `
		resource "google_organization_iam_binding" "binding" {
			org_id = data.google_organization.org.id
			role    = "roles/browser"
			
			members = [
				"user:alice@gmail.com",
			]
		}`,
			expected: iam.Binding{
				Metadata: iacTypes.NewTestMetadata(),
				Members: []iacTypes.StringValue{
					iacTypes.String("user:alice@gmail.com", iacTypes.NewTestMetadata())},
				Role:                          iacTypes.String("roles/browser", iacTypes.NewTestMetadata()),
				IncludesDefaultServiceAccount: iacTypes.Bool(false, iacTypes.NewTestMetadata()),
			},
		},
		{
			name: "defaults",
			terraform: `
		resource "google_organization_iam_binding" "binding" {
		}`,
			expected: iam.Binding{
				Metadata:                      iacTypes.NewTestMetadata(),
				Role:                          iacTypes.String("", iacTypes.NewTestMetadata()),
				IncludesDefaultServiceAccount: iacTypes.Bool(false, iacTypes.NewTestMetadata()),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			modules := tftestutil.CreateModulesFromSource(t, test.terraform, ".tf")
			adapted := AdaptBinding(modules.GetBlocks()[0], modules)
			testutil.AssertDefsecEqual(t, test.expected, adapted)
		})
	}
}

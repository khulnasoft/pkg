package computing

import (
	"testing"

	"go.khulnasoft.com/tunnel/internal/testutil"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/tftestutil"
	"go.khulnasoft.com/pkg/iac/providers/nifcloud/computing"
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

func Test_adaptSecurityGroups(t *testing.T) {
	tests := []struct {
		name      string
		terraform string
		expected  []computing.SecurityGroup
	}{
		{
			name: "configured",
			terraform: `
			resource "nifcloud_security_group" "example" {
				group_name = "example"
				description = "memo"
			}
			
			resource "nifcloud_security_group_rule" "example" {
				type                 = "IN"
				security_group_names = [nifcloud_security_group.example.group_name]
				from_port            = 22
				to_port              = 22
				protocol             = "TCP"
				description          = "memo"
				cidr_ip              = "1.2.3.4/32"
			}
`,
			expected: []computing.SecurityGroup{{
				Metadata:    iacTypes.NewTestMetadata(),
				Description: iacTypes.String("memo", iacTypes.NewTestMetadata()),
				IngressRules: []computing.SecurityGroupRule{
					{
						Metadata:    iacTypes.NewTestMetadata(),
						CIDR:        iacTypes.String("1.2.3.4/32", iacTypes.NewTestMetadata()),
						Description: iacTypes.String("memo", iacTypes.NewTestMetadata()),
					},
				},
			}},
		},
		{
			name: "defaults",
			terraform: `
			resource "nifcloud_security_group" "example" {
			}
			
			resource "nifcloud_security_group_rule" "example" {
				type                 = "IN"
				security_group_names = [nifcloud_security_group.example.group_name]
			}

`,

			expected: []computing.SecurityGroup{{
				Metadata:    iacTypes.NewTestMetadata(),
				Description: iacTypes.String("", iacTypes.NewTestMetadata()),
				IngressRules: []computing.SecurityGroupRule{
					{
						Metadata:    iacTypes.NewTestMetadata(),
						CIDR:        iacTypes.String("", iacTypes.NewTestMetadata()),
						Description: iacTypes.String("", iacTypes.NewTestMetadata()),
					},
				},
			}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			modules := tftestutil.CreateModulesFromSource(t, test.terraform, ".tf")
			sgAdapter := sgAdapter{sgRuleIDs: modules.GetChildResourceIDMapByType("nifcloud_security_group_rule")}
			adapted := sgAdapter.adaptSecurityGroups(modules)
			testutil.AssertDefsecEqual(t, test.expected, adapted)
		})
	}
}

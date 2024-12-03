package dns

import (
	"testing"

	"go.khulnasoft.com/tunnel/internal/testutil"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/tftestutil"
	"go.khulnasoft.com/pkg/iac/providers/nifcloud/dns"
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

func Test_adaptRecords(t *testing.T) {
	tests := []struct {
		name      string
		terraform string
		expected  []dns.Record
	}{
		{
			name: "configured",
			terraform: `
			resource "nifcloud_dns_record" "example" {
				type    = "A"
				record  = "example-record"
			}
`,
			expected: []dns.Record{{
				Metadata: iacTypes.NewTestMetadata(),
				Type:     iacTypes.String("A", iacTypes.NewTestMetadata()),
				Record:   iacTypes.String("example-record", iacTypes.NewTestMetadata()),
			}},
		},
		{
			name: "defaults",
			terraform: `
			resource "nifcloud_dns_record" "example" {
			}
`,

			expected: []dns.Record{{
				Metadata: iacTypes.NewTestMetadata(),
				Type:     iacTypes.String("", iacTypes.NewTestMetadata()),
				Record:   iacTypes.String("", iacTypes.NewTestMetadata()),
			}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			modules := tftestutil.CreateModulesFromSource(t, test.terraform, ".tf")
			adapted := adaptRecords(modules)
			testutil.AssertDefsecEqual(t, test.expected, adapted)
		})
	}
}

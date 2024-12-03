package rdb

import (
	"testing"

	"go.khulnasoft.com/tunnel/internal/testutil"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/tftestutil"
	"go.khulnasoft.com/pkg/iac/providers/nifcloud/rdb"
	iacTypes "go.khulnasoft.com/pkg/iac/types"
)

func Test_adaptDBInstances(t *testing.T) {
	tests := []struct {
		name      string
		terraform string
		expected  []rdb.DBInstance
	}{
		{
			name: "configured",
			terraform: `
			resource "nifcloud_db_instance" "example" {
				backup_retention_period = 2
				engine                  = "MySQL"
				engine_version          = "5.7.15"
				publicly_accessible     = false
				network_id              = "example-network"
			}
`,
			expected: []rdb.DBInstance{{
				Metadata:                  iacTypes.NewTestMetadata(),
				BackupRetentionPeriodDays: iacTypes.Int(2, iacTypes.NewTestMetadata()),
				Engine:                    iacTypes.String("MySQL", iacTypes.NewTestMetadata()),
				EngineVersion:             iacTypes.String("5.7.15", iacTypes.NewTestMetadata()),
				NetworkID:                 iacTypes.String("example-network", iacTypes.NewTestMetadata()),
				PublicAccess:              iacTypes.Bool(false, iacTypes.NewTestMetadata()),
			}},
		},
		{
			name: "defaults",
			terraform: `
			resource "nifcloud_db_instance" "example" {
			}
`,

			expected: []rdb.DBInstance{{
				Metadata:                  iacTypes.NewTestMetadata(),
				BackupRetentionPeriodDays: iacTypes.Int(0, iacTypes.NewTestMetadata()),
				Engine:                    iacTypes.String("", iacTypes.NewTestMetadata()),
				EngineVersion:             iacTypes.String("", iacTypes.NewTestMetadata()),
				NetworkID:                 iacTypes.String("net-COMMON_PRIVATE", iacTypes.NewTestMetadata()),
				PublicAccess:              iacTypes.Bool(true, iacTypes.NewTestMetadata()),
			}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			modules := tftestutil.CreateModulesFromSource(t, test.terraform, ".tf")
			adapted := adaptDBInstances(modules)
			testutil.AssertDefsecEqual(t, test.expected, adapted)
		})
	}
}

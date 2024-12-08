// © 2022-2023 Khulnasoft Limited All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/khulnasoft/policy-engine/pkg/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version",
	RunE: func(cmd *cobra.Command, args []string) error {
		v := version.GetVersionInfo()

		table := [][2]string{}
		table = append(table, [2]string{"Version", v.Version})
		table = append(table, [2]string{"OPA Version", v.OPAVersion})
		table = append(table, [2]string{"Terraform Version", v.TerraformVersion})
		revision := v.Revision
		if v.HasChanges {
			revision = fmt.Sprintf("%s*", revision)
		}
		table = append(table, [2]string{"Revision", revision})

		padding := 0
		for _, row := range table {
			if len(row[0]) > padding {
				padding = len(row[0])
			}
		}

		for _, row := range table {
			fmt.Fprintf(os.Stdout, "%s:%s%s\n",
				row[0],
				strings.Repeat(" ", padding-len(row[0])+1),
				row[1],
			)
		}

		return nil
	},
}

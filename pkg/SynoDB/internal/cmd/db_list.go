package cmd

import (
	"sort"

	"github.com/spf13/cobra"
	"github.com/khulnasoft/synodb/internal/synodb"
	"golang.org/x/exp/slices"
)

func init() {
	dbCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:               "list",
	Short:             "List databases.",
	Args:              cobra.NoArgs,
	ValidArgsFunction: noFilesArg,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		client, err := createSynodbClientFromAccessToken(true)
		if err != nil {
			return err
		}

		databases, err := client.Databases.List()
		if err != nil {
			return err
		}
		setDatabasesCache(databases)

		printDBListTable(databases)
		return nil
	},
}

func printDBListTable(databases []synodb.Database) {
	headers, data := dbListTable(databases)
	if !shouldPrintLocations(databases) {
		headers, data = removeColumn(headers, data, "Locations")
	}
	if !shouldPrintGroups(databases) {
		headers, data = removeColumn(headers, data, "Group")
	}
	printTable(headers, data)
}

func shouldPrintLocations(databases []synodb.Database) bool {
	for _, database := range databases {
		if database.Group == "" {
			return true
		}
	}
	return false
}

func shouldPrintGroups(databases []synodb.Database) bool {
	mp := map[string]bool{}
	for _, database := range databases {
		mp[database.Group] = true
	}
	return len(mp) > 1
}

func dbListTable(databases []synodb.Database) (headers []string, data [][]string) {
	for _, database := range databases {
		row := []string{database.Name, getDatabaseLocations(database), formatGroup(database.Group), getDatabaseUrl(&database)}
		data = append(data, row)
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})

	return []string{"Name", "Locations", "Group", "URL"}, data
}

func removeColumn(headers []string, data [][]string, column string) ([]string, [][]string) {
	i := slices.Index(headers, column)
	if i == -1 {
		return headers, data
	}

	filtered := make([][]string, 0, len(data))
	for _, row := range data {
		filtered = append(filtered, removeIndex(row, i))
	}

	return removeIndex(headers, i), filtered
}

func removeIndex(arr []string, i int) []string {
	return append(arr[:i], arr[i+1:]...)
}

func formatGroup(group string) string {
	if group == "" {
		return "-"
	}
	return group
}

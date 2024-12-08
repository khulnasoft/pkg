package cmd

import (
	"github.com/spf13/cobra"
	"github.com/khulnasoft/synodb/internal"
)

func init() {
	apiTokensCmd.AddCommand(listApiTokensCmd)
}

var listApiTokensCmd = &cobra.Command{
	Use:   "list",
	Short: "List API tokens.",
	Long: "" +
		"API tokens are revocable non-expiring tokens that authenticate holders as the user who minted them.\n" +
		"They can be used to implement automations with the " + internal.Emph("synodb") + " CLI or the platform API.",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true

		client, err := createSynodbClientFromAccessToken(true)
		if err != nil {
			return err
		}

		apiTokens, err := client.ApiTokens.List()
		if err != nil {
			return err
		}

		data := [][]string{}
		for _, apiToken := range apiTokens {
			data = append(data, []string{apiToken.Name})
		}
		printTable([]string{"Name"}, data)

		return nil
	},
}

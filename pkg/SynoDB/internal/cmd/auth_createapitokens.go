package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/khulnasoft/synodb/internal"
	"github.com/khulnasoft/synodb/internal/prompt"
	"github.com/khulnasoft/synodb/internal/synodb"
)

func init() {
	apiTokensCmd.AddCommand(createApiTokensCmd)
}

var createApiTokensCmd = &cobra.Command{
	Use:   "mint api_token_name",
	Short: "Mint an API token.",
	Long: "" +
		"API tokens are revocable non-expiring tokens that authenticate holders as the user who minted them.\n" +
		"They can be used to implement automations with the " + internal.Emph("synodb") + " CLI or the platform API.",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true

		client, err := createSynodbClientFromAccessToken(true)
		if err != nil {
			return err
		}

		name := strings.TrimSpace(args[0])

		if err := synodb.CheckName(name); err != nil {
			return fmt.Errorf("invalid token name: %w", err)
		}

		description := fmt.Sprintf("Creating api token %s", internal.Emph(name))
		bar := prompt.Spinner(description)
		defer bar.Stop()

		data, err := client.ApiTokens.Create(name)
		if err != nil {
			return err
		}

		bar.Stop()
		fmt.Println(data.Value)
		return nil
	},
}

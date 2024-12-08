package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/khulnasoft/synodb/internal"
)

func init() {
	rootCmd.AddCommand(quickstartCmd)
}

var quickstartCmd = &cobra.Command{
	Use:               "quickstart",
	Short:             "Synodb quick quickstart.",
	Args:              cobra.NoArgs,
	ValidArgsFunction: noFilesArg,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		fmt.Print("\nWelcome to Synodb!\n\n")
		fmt.Printf("If you are a new user, please sign up with %s; otherwise login\n", internal.Emph("synodb auth signup"))
		fmt.Printf("with %s. When you are authenticated, you can create a new\n", internal.Emph("synodb auth login"))
		fmt.Printf("database with %s. You can also run %s for help.\n", internal.Emph("synodb db create"), internal.Emph("synodb help"))
		fmt.Printf("\nFor a more comprehensive getting started guide, open the following URL:\n\n")
		fmt.Printf("  https://docs.synodb.tech/tutorials/get-started-synodb\n\n")
		return nil
	},
}

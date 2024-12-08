package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/khulnasoft/synodb/internal"
)

var enableExtensionsFlag bool

func addEnableExtensionsFlag(cmd *cobra.Command) {
	usage := []string{
		"Enables experimental support for SQLite extensions.",
		"If you would like to see some extension included, run " + internal.Emph("synodb account feedback") + ".",
		internal.Warn("Warning") + ": extensions support is experimental and subject to change",
	}
	cmd.Flags().BoolVar(&enableExtensionsFlag, "enable-extensions", false, strings.Join(usage, "\n"))
}

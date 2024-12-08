package cmd

import (
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var accountBookMeetingCmd = &cobra.Command{
	Use:               "bookmeeting",
	Short:             "Book a meeting with the Synodb team.",
	Args:              cobra.NoArgs,
	ValidArgsFunction: noFilesArg,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		return browser.OpenURL("https://calendly.com/d/gt7-bfd-83n/meet-with-chiselstrike")
	},
}

package cron

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "cron",
		RunE: runCron,
	}

	cmd.AddCommand(
		newSectionCommand(),
	)

	return cmd
}

func runCron(_ *cobra.Command, _ []string) error {
	return nil
}

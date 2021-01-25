package feature

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "feature",
		Short: "run feature examples",
	}

	cmd.AddCommand()
	return cmd
}

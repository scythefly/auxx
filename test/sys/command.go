package sys

import (
	"github.com/spf13/cobra"

	"auxx/test/sys/flock"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sys",
		Short: "run sys examples",
	}

	cmd.AddCommand(
		flock.NewCommand(),
	)

	return cmd
}

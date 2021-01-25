package storage

import (
	"github.com/spf13/cobra"

	"auxx/test/storage/s3"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "storage",
		Short: "run storage examples",
	}

	cmd.AddCommand(
		s3.NewCommand(),
	)
	return cmd
}

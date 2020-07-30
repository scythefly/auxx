package sync

import (
	"github.com/spf13/cobra"
)

func newAtomicCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "atomic",
		Short: "Run sync atomic examples",
		RunE:  syncAtomicRun,
	}

	return cmd
}

func syncAtomicRun(_ *cobra.Command, _ []string) error {
	var err error

	return err
}

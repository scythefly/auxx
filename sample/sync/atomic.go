package sync

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/atomic"
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

	var atom atomic.Uint32
	atom.Store(42)
	fmt.Println(atom.Inc())
	fmt.Println(atom.CAS(43, 10))
	fmt.Println(atom.Load())

	return err
}

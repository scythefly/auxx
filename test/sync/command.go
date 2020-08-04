package sync

import (
	"github.com/spf13/cobra"

	"auxx/types"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "Run sync examples",
		RunE:  syncRun,
	}

	cmd.AddCommand(
		newChanCommand(),
		newAtomicCommand(),
	)
	return cmd
}

type syncStruct struct {
	str string
}

func syncRun(*cobra.Command, []string) error {
	var err error

	var s syncStruct
	types.G.Go(func() error {
		syncGo(&s)
		return nil
	})
	types.G.Go(func() error {
		syncGo(&s)
		return nil
	})

	types.G.Wait()
	return err
}

func syncGo(s *syncStruct) {
	var cnt int
	for cnt < 10000 {
		if s.str == "" {
			s.str = "starting"
		} else if s.str == "running" {
			s.str = "starting"
		}
		cnt++
	}
}

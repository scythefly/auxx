package sample

import (
	"github.com/spf13/cobra"

	"auxx/sample/cron"
	"auxx/sample/encoding"
	"auxx/sample/feature"
	"auxx/sample/image"
	"auxx/sample/log"
	"auxx/sample/lru"
	"auxx/sample/net"
	"auxx/sample/storage"
	"auxx/sample/sync"
	"auxx/sample/sys"
	"auxx/sample/update"
)

func NewTestCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test",
		Short: `Run test examples`,
	}

	cmd.AddCommand(
		cron.NewCommand(),
		encoding.NewCommand(),
		feature.NewCommand(),
		image.NewCommand(),
		log.NewCommand(),
		lru.NewCommand(),
		net.NewCommand(),
		storage.NewCommand(),
		sync.NewCommand(),
		sys.NewCommand(),
		update.NewCommand(),
	)
	return cmd
}

package command

import (
	"github.com/spf13/cobra"

	"auxx/test/cron"
	"auxx/test/daemon"
	"auxx/test/encoding"
	"auxx/test/etcd"
	"auxx/test/feature"
	"auxx/test/image"
	"auxx/test/kafka"
	"auxx/test/log"
	"auxx/test/lru"
	"auxx/test/net"
	"auxx/test/storage"
	"auxx/test/sync"
	"auxx/test/sys"
	"auxx/test/update"
)

func newTestCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test",
		Short: `Run test examples`,
	}

	cmd.AddCommand(
		cron.NewCommand(),
		daemon.NewCommand(),
		encoding.NewCommand(),
		etcd.NewCommand(),
		feature.NewCommand(),
		image.NewCommand(),
		kafka.NewCommand(),
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

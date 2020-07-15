package etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	endpoints = []string{"10.68.192.112:2379", "10.68.192.113:2379", "10.68.192.114:2379"}
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "etcd",
		Short: "Run etcd examples",
		RunE:  etcdRun,
	}

	cmd.AddCommand(
		newWatchCommand(),
		newPluginCommand(),
	)
	return cmd
}

func etcdRun(cmd *cobra.Command, args []string) error {
	// put get
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return errors.WithMessagef(err, "etcd run")
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	resp, err := cli.Put(ctx, "sample_key", time.Now().Format(time.RFC3339Nano))
	cancel()
	if err != nil {
		return errors.WithMessagef(err, "etcd run")
	}
	fmt.Println(resp.Header.ClusterId)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	vresp, err := cli.Get(ctx, "sample_key")
	if err != nil {
		return errors.WithMessagef(err, "etcd run")
	}
	for _, kvs := range vresp.Kvs {
		fmt.Println(string(kvs.Key), string(kvs.Value))
	}

	return nil
}

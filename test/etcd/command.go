package etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var (
	endpoints = []string{"127.0.0.2:2379"}
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
		newLeaseCommand(),
		newLeaseKVCommand(),
		newCampaignCommand(),
		newKvCommand(),
	)
	return cmd
}

type rootC struct {
	cli  *clientv3.Client
	g    *errgroup.Group
	gctx context.Context
}

func etcdRun(cmd *cobra.Command, args []string) error {
	var err error
	rc := &rootC{}
	// put get
	rc.cli, err = clientv3.New(clientv3.Config{
		Endpoints:            endpoints,
		DialTimeout:          5 * time.Second,
		DialKeepAliveTime:    3 * time.Second,
		DialKeepAliveTimeout: 2 * time.Second,
	})
	if err != nil {
		return errors.WithMessagef(err, "etcd run")
	}
	defer rc.cli.Close()

	fmt.Println(time.Now().Unix())
	rc.g, rc.gctx = errgroup.WithContext(context.Background())
	rc.g.Go(rc.put)
	rc.g.Go(rc.get)
	rc.g.Go(func() error {
		select {
		case <-rc.gctx.Done():
			return errors.New("rc.gctx done")
		case <-rc.cli.Ctx().Done():
			return errors.New("rc.cli.Ctx.Done")
		}
	})

	fmt.Println(rc.g.Wait())
	fmt.Println(time.Now().Unix())

	return nil
}

func (r *rootC) get() error {
	ctx, _ := context.WithTimeout(r.gctx, 8*time.Second)
	vresp, err := r.cli.Get(ctx, "sample_key")
	if err != nil {
		return errors.WithMessagef(err, "etcd run")
	}
	fmt.Println(">>> resp kvs len: ", len(vresp.Kvs))
	for _, kvs := range vresp.Kvs {
		fmt.Println(string(kvs.Key), string(kvs.Value))
	}
	return err
}

func (r *rootC) put() error {
	ctx, cancel := context.WithTimeout(r.gctx, 10*time.Second)
	resp, err := r.cli.Put(ctx, "sample_key", time.Now().Format(time.RFC3339Nano))
	cancel()
	if err != nil {
		return errors.WithMessagef(err, "etcd run")
	}
	fmt.Println(resp.Header.ClusterId)
	return err
}

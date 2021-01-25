package etcd

import (
	"context"
	"fmt"
	"time"

	v3 "github.com/coreos/etcd/clientv3"
	"github.com/spf13/cobra"
	atomic2 "go.uber.org/atomic"
	"golang.org/x/sync/errgroup"
)

var (
	_endpoints  = []string{"36.99.187.199:22379"}
	_endpoints1 = []string{"36.99.187.196:22379"}
	_cnt        atomic2.Int64
)

const (
	_Puters   = 1000
	_Watchers = 1000
)

func newWatch1Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "watch1",
		Short: "Run etcd watch1 examples",
		RunE:  etcdWatch1,
	}
	return cmd
}

func etcdWatch1(*cobra.Command, []string) error {
	var err error
	ctx := context.Background()
	cli, err := v3.New(v3.Config{
		Endpoints:   _endpoints,
		DialTimeout: 5 * time.Second,
		Context:     ctx,
	})
	if err != nil {
		return err
	}
	cli1, err := v3.New(v3.Config{
		Endpoints:   _endpoints1,
		DialTimeout: 5 * time.Second,
		Context:     ctx,
	})
	if err != nil {
		return err
	}
	cctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	if _, err = cli.Get(cctx, "check_server_available"); err != nil {
		return err
	}
	if _, err = cli1.Get(cctx, "check_server_available"); err != nil {
		return err
	}
	cancel()

	var g errgroup.Group
	for i := 0; i < _Watchers; i++ {
		g.Go(func() error {
			return goWatch(cli1)
		})
	}
	time.Sleep(time.Second)
	for i := 0; i < _Puters; i++ {
		idx := i
		g.Go(func() error {
			goPut(cli, idx)
			return nil
		})
	}
	g.Go(goObserve)

	g.Wait()
	return nil
}

func goPut(cli *v3.Client, idx int) {
	format := fmt.Sprintf("aaa_%d_%%d", idx)
	for i := 0; i < 100; i++ {
		str := fmt.Sprintf(format, i)
		cli.Put(context.Background(), str, str)
	}
}

func goWatch(cli *v3.Client) error {
	// var modRev int64
	// resp, _ := cli.Get(context.Background(), "aaa_", v3.WithPrefix())
	// for _, kvs := range resp.Kvs {
	// 	_cnt.Inc()
	// 	if kvs.ModRevision > modRev {
	// 		modRev = kvs.ModRevision
	// 	}
	// }

	watcher := v3.NewWatcher(cli)
	wcc := watcher.Watch(context.Background(), "aaa_", v3.WithPrefix() /*, v3.WithRev(modRev)*/)
	for {
		select {
		case resp := <-wcc:
			if resp.Canceled {
				return resp.Err()
			}
			for range resp.Events {
				_cnt.Inc()
			}
		}
	}
}

func goObserve() error {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println(_cnt.Load() / _Watchers)
		}
	}
}

package etcd

import (
	"plugin"
	"reflect"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// type ServeFunc func(*clientv3.Client) error
type Service interface {
	Serve(*clientv3.Client) error
}

func newPluginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plugin",
		Short: "Run etcd plugin examples",
		RunE:  etcdPlugin,
	}

	return cmd
}

func etcdPlugin(*cobra.Command, []string) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return errors.WithMessage(err, "etcd watch")
	}
	defer cli.Close()
	p, err := plugin.Open("./plugin_etcd_client.so")
	if err != nil {
		return errors.WithMessage(err, "etcd plugin open")
	}
	cc, err := p.Lookup("EtcdClient")
	if err != nil {
		return errors.WithMessage(err, "etcd plugin load \"Serve\"")
	}
	c, ok := cc.(Service)
	if !ok {
		tt := reflect.TypeOf(cc)
		return errors.Errorf("unexpected type from plugin symbol: %s", tt.Name())
	}

	gW.Go(func() error {
		return c.Serve(cli)
	})

	gP.Go(func() error {
		return produce(1)
	})
	gP.Go(func() error {
		return produce(2)
	})
	gP.Wait()
	return nil
}

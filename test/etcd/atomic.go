package etcd

import (
	"context"
	"fmt"
	"path/filepath"
	"sync"
	"time"

	v3 "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
	"pkg/util"
)

func newAtomicCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "atomic",
		Short: "Run atomic examples",
		RunE:  atomicRun,
	}

	return cmd
}

func atomicRun(_ *cobra.Command, _ []string) error {
	var a atomic
	return a.run()
}

type atomic struct {
	cli       *v3.Client
	kv        sync.Map
	keyPrefix string
}

func (a *atomic) run() error {
	var err error
	a.keyPrefix = "etcd/atomic/"
	a.cli, err = v3.New(v3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return errors.WithMessagef(err, "etcd run")
	}
	defer a.cli.Close()

	var g errgroup.Group

	g.Go(a.watch)
	g.Go(a.fake)

	return g.Wait()
}

func (a *atomic) watch() (err error) {
	cli := a.cli
	resp, err := cli.Get(context.Background(), a.keyPrefix, v3.WithPrefix())
	if err != nil {
		fmt.Println("watcher, get alive info err", err)
		return
	}
	for _, kvs := range resp.Kvs {
		a.kv.Store(string(kvs.Key), string(kvs.Value))
	}

	watcher := v3.NewWatcher(cli)
	wcc := watcher.Watch(context.Background(), a.keyPrefix, v3.WithPrefix())

	for {
		select {
		case resp := <-wcc:
			for _, event := range resp.Events {
				if event.IsCreate() {
					a.kv.Store(string(event.Kv.Key), string(event.Kv.Value))
				} else if event.Type == mvccpb.DELETE {
					a.kv.Delete(string(event.Kv.Key))
				}
			}
		}
	}
}

func (a *atomic) fake() error {
	ticker := time.NewTicker(time.Second)
	var cnt int
	for {
		select {
		case <-ticker.C:
			cnt++
			key := filepath.Join(a.keyPrefix, util.RandomString(5))
			value := util.RandomString(6)
			a.cli.Put(context.Background(), key, value)

			if cnt%10 == 0 {
				a.cli.Put(context.Background(), key, "xxxxxxx")
				go a.get(key)
				go a.kvGet(key)
			}
		}
	}
}

func (a *atomic) kvGet(key string) {
	if v, ok := a.kv.Load(key); ok {
		fmt.Println("kvGet:", v.(string))
	}
}

func (a *atomic) get(key string) {
	resp, _ := a.cli.Get(context.Background(), key)
	for _, kvs := range resp.Kvs {
		fmt.Println(">> get:", string(kvs.Value))
	}
}

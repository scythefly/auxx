package etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
	set "github.com/deckarep/golang-set"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	leaseAgent set.Set
)

func newLeaseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lease",
		Short: "Run etcd lease examples",
		RunE:  etcdLease,
	}
	return cmd
}

func etcdLease(*cobra.Command, []string) error {
	leaseAgent = set.NewSet()
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return errors.WithMessagef(err, "etcd run")
	}
	defer cli.Close()

	for i := 0; i < 5; i++ {
		idx := i
		go leaseAliveClient(idx)
	}

	time.Sleep(5 * time.Second)

	go watchAlive(cli)

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			var cnt int
			leaseAgent.Each(func(interface{}) bool {
				cnt++
				return false
			})
			fmt.Println("core alive:", cnt)
		}
	}
}

func watchAlive(cli *clientv3.Client) {
	resp, err := cli.Get(context.Background(), "core/alive/", clientv3.WithPrefix())
	if err != nil {
		fmt.Println("watcher, get alive info err", err)
		return
	}
	for _, kvs := range resp.Kvs {
		leaseAgent.Add(string(kvs.Key))
	}

	watcher := clientv3.NewWatcher(cli)
	wcc := watcher.Watch(context.Background(), "core/alive/", clientv3.WithPrefix())

	for resp := range wcc {
		for _, event := range resp.Events {
			fmt.Println("IsCreate:", event.IsCreate(), "IsModify:", event.IsModify())
			fmt.Println(string(event.Kv.Key), string(event.Kv.Value))
			if event.PrevKv != nil {
				fmt.Println("PrevKv", string(event.PrevKv.Key), string(event.PrevKv.Value))
			}
			if string(event.Kv.Value) == "alive" {
				leaseAgent.Add(string(event.Kv.Key))
			} else {
				leaseAgent.Remove(string(event.Kv.Key))
			}
		}
	}
}

func leaseAliveClient(idx int) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(idx, "new client failed", err)
		return
	}
	defer cli.Close()
	alive, err := cli.Grant(context.TODO(), 3)
	if err != nil {
		fmt.Println(idx, "create lease key failed")
		return
	}
	_, err = cli.Put(context.TODO(), fmt.Sprintf("core/alive/%d", idx), "alive", clientv3.WithLease(alive.ID))
	if err != nil {
		fmt.Println(idx, "put lease alive key err", err)
		return
	}
	ch, kaerr := cli.KeepAlive(context.TODO(), alive.ID)
	if kaerr != nil {
		fmt.Println(idx, "keep alive err", kaerr)
		return
	}

	var cnt int
	for range ch {
		cnt++
		if cnt > (idx+1)*6 {
			break
		}
	}
	fmt.Println(idx, "quit")
}

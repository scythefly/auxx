package etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func newLeaseKVCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lease_kv",
		Short: "Run etcd lease examples",
		RunE:  etcdLeaseKV,
	}
	return cmd
}

func etcdLeaseKV(*cobra.Command, []string) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return errors.WithMessagef(err, "etcd run")
	}
	defer cli.Close()

	go watchLeaseKV(cli)

	key := "core/lease_kv/" + time.Now().Format(time.RFC3339)
	lkv, err := cli.Grant(context.TODO(), 10)
	if err != nil {
		return errors.WithMessage(err, "create lease key")
	}
	_, err = cli.Put(context.TODO(), key, "123", clientv3.WithLease(lkv.ID))
	if err != nil {
		return errors.WithMessage(err, "put lease key")
	}
	go keepAlive(cli, lkv.ID, key, "123")

	time.Sleep(3 * time.Second)
	lkv1, err := cli.Grant(context.TODO(), 10)
	if err != nil {
		return errors.WithMessage(err, "create lease key")
	}
	_, err = cli.Put(context.TODO(), key, "456", clientv3.WithLease(lkv1.ID))
	if err != nil {
		return errors.WithMessage(err, "put lease key")
	}
	go keepAlive(cli, lkv1.ID, key, "456")

	select {}
}

func watchLeaseKV(cli *clientv3.Client) {
	watcher := clientv3.NewWatcher(cli)
	wcc := watcher.Watch(context.Background(), "core/lease_kv/", clientv3.WithPrefix())

	for resp := range wcc {
		for _, event := range resp.Events {
			fmt.Println("IsCreate:", event.IsCreate(), "IsModify:", event.IsModify())
			fmt.Println(string(event.Kv.Key), string(event.Kv.Value))
			if event.PrevKv != nil {
				fmt.Println("PrevKv", string(event.PrevKv.Key), string(event.PrevKv.Value))
			}
		}
	}
}

func keepAlive(cli *clientv3.Client, id clientv3.LeaseID, key, value string) {
	var cnt int
	for {
		cnt++
		time.Sleep(5 * time.Second)
		resp, err := cli.Get(context.TODO(), key)
		if err != nil {
			fmt.Println(id, "get:", err)
			return
		}
		var respValue string
		for _, kvs := range resp.Kvs {
			respValue = string(kvs.Value)
			fmt.Println("value -- ", respValue)
		}
		if value != respValue {
			if cnt < 3 {
				continue
			}
			fmt.Println("reset, break")
		}

		ka, kaerr := cli.KeepAliveOnce(context.TODO(), id)
		if kaerr != nil {
			fmt.Println(id, "keep alive err", kaerr)
			return
		}
		fmt.Println(id, "ttl:", ka.TTL)
	}
}

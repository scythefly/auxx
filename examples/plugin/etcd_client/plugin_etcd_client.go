package main

// go build -buildmode=plugin -o plugin_etcd_client.so plugin_etcd_client.go

import (
	"context"
	"fmt"
	"runtime"

	"github.com/coreos/etcd/clientv3"

	"auxx/examples/plugin/etcd_client/version"
)

func init() {
	fmt.Printf(`plugin etcd_client init
	Version: %s
	build with: %s, at %s
`, version.Version, runtime.Version(), version.Built)
}

type Client struct{}

// exported Serve
func (c *Client) Serve(cli *clientv3.Client) error {
	watcher := clientv3.NewWatcher(cli)
	wcc := watcher.Watch(context.Background(), "scythefly.top/", clientv3.WithPrefix())

	var cnt int
	for resp := range wcc {
		cnt++
		fmt.Printf("Watcher - %d >>> \n", cnt)
		for _, event := range resp.Events {
			fmt.Println("IsCreate:", event.IsCreate(), "IsModify:", event.IsModify())
			fmt.Println(string(event.Kv.Key), string(event.Kv.Value))
			if event.PrevKv != nil {
				fmt.Println("PrevKv", string(event.PrevKv.Key), string(event.PrevKv.Value))
			}
		}
	}
	return nil
}

// export EtcdClient
var EtcdClient = Client{}

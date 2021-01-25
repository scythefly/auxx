package etcd_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	v3 "github.com/coreos/etcd/clientv3"
)

var _endpoints = []string{"10.68.2.112:2379"}

func Test_Get(t *testing.T) {
	var err error
	ctx := context.Background()
	cli, err := v3.New(v3.Config{
		Endpoints:   _endpoints,
		DialTimeout: 5 * time.Second,
		Context:     ctx,
	})
	if err != nil {
		t.Error(err)
		return
	}
	cctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	if _, err = cli.Get(cctx, "check_server_available"); err != nil {
		t.Error(err)
		return
	}
	cancel()

	t1 := time.Now().UnixNano() / 1e6
	for i := 0; i < 10000; i++ {
		cli.Put(ctx, fmt.Sprintf("aaa_%d", i), fmt.Sprintf("bbb_%d", i))
	}
	t2 := time.Now().UnixNano() / 1e6
	t.Log("put 10000 keys, cost:", t2-t1, "ms")

	t1 = time.Now().UnixNano() / 1e6
	resp, _ := cli.Get(ctx, "aaa_", v3.WithPrefix())
	var modRev, cnt int64
	for _, kvs := range resp.Kvs {
		cnt++
		if kvs.ModRevision > modRev {
			modRev = kvs.ModRevision
		}
	}
	t2 = time.Now().UnixNano() / 1e6
	t.Log("get aaa_ with prefix, keys:", cnt, "cost", t2-t1, "ms, max mod rev:", modRev)
}

package etcd_test

import (
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"
)

var (
	endpoints = []string{"10.68.192.112:2379", "10.68.192.113:2379", "10.68.192.114:2379"}
)

func Test_LeaseKeepAlive(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		t.Error(errors.WithMessage(err, "etcd watch"))
	}
	defer cli.Close()
}

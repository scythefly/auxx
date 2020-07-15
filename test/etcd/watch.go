package etcd

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var (
	mu sync.Mutex
	gW errgroup.Group
	gP errgroup.Group
)

func newWatchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "watch",
		Short: "Run etcd watch examples",
		RunE:  etcdWatch,
	}
	return cmd
}

func etcdWatch(*cobra.Command, []string) error {
	var err error
	stopped := make(chan bool)

	gW.Go(func() error {
		return watch(1, stopped)
	})

	gW.Go(func() error {
		return watch(2, stopped)
	})

	gW.Go(func() error {
		return watch(2, stopped)
	})

	gP.Go(func() error {
		return produce(1)
	})
	gP.Go(func() error {
		return produce(2)
	})
	gP.Go(func() error {
		return produce(3)
	})
	gP.Wait()

	close(stopped)

	gW.Wait()
	return err
}

func watch(id int, stopped chan bool) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return errors.WithMessage(err, "etcd watch")
	}
	defer cli.Close()

	watcher := clientv3.NewWatcher(cli)
	wcc := watcher.Watch(context.Background(), "scythefly.top/", clientv3.WithPrefix())

	var cnt int
	for {
		cnt++
		select {
		case resp := <-wcc:
			mu.Lock()
			fmt.Printf("Watcher[%d] - %d >>> \n", id, cnt)
			for _, event := range resp.Events {
				fmt.Println("IsCreate:", event.IsCreate(), "IsModify:", event.IsModify())
				fmt.Println(string(event.Kv.Key), string(event.Kv.Value))
				if event.PrevKv != nil {
					fmt.Println("PrevKv", string(event.PrevKv.Key), string(event.PrevKv.Value))
				}
			}
			mu.Unlock()
		case <-stopped:
			goto END
		}
	}
END:
	return nil
}

func produce(id int) error {
	var err error
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return errors.WithMessage(err, "etcd watch")
	}
	ticker := time.NewTicker(time.Second)
	var cnt int
	for {
		select {
		case <-ticker.C:
			cnt++
			if cnt > 100 {
				return nil
			}
			key := fmt.Sprintf("scythefly.top/%d", cnt%10)
			// fmt.Printf(">>>>>> Producer[%d] put %s\n", id, key)
			if _, err = cli.Put(context.Background(), key, time.Now().Format(time.RFC3339Nano)); err != nil {
				fmt.Println("put routine err:", err.Error())
				return err
			}
		}
	}
}

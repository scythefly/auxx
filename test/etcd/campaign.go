package etcd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/spf13/cobra"
)

func newCampaignCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "campaign",
		Short: "Run campaign examples",
		RunE:  campaignRun,
	}

	return cmd
}

func campaignRun(_ *cobra.Command, _ []string) error {
	var err error
	endpoints := []string{"127.0.0.1:2379"}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	go campaign(cli)

	select {}
	// return err
}

func campaign(cli *clientv3.Client) {
	for {
		fmt.Println(">>>>> 1")
		s, err := concurrency.NewSession(cli, concurrency.WithTTL(15))
		fmt.Println(">>>>> 11")
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(">>>>> 2")
		e := concurrency.NewElection(s, "campaign_test")
		ctx := context.TODO()

		fmt.Println(">>>>> 3")
		if err = e.Campaign(ctx, "val"); err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("elect: success")
		dodo(s)
	}
}

func dodo(s *concurrency.Session) {
	for {
		select {
		case <-s.Done():
			return
		default:
		}
		fmt.Println("I am the leader!!!")
		time.Sleep(time.Second)
	}
}

func doCrontab() {
}

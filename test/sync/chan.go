package sync

import (
	"fmt"
	"strconv"
	"time"

	"github.com/scythefly/orb"
	"github.com/spf13/cobra"

	"auxx/types"
)

var (
	upchan  = make(chan []byte, 12)
	clients = orb.NewSet()
)

func newChanCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chan",
		Short: "Run chan examples",
		RunE:  chanRun,
	}

	cmd.AddCommand(
		newChanCloseCommand(),
	)
	return cmd
}

func chanRun(_ *cobra.Command, args []string) error {
	var err error
	var cs = 1500
	if len(args) > 0 {
		cs, _ = strconv.Atoi(args[0])
	}
	chanTest(cs)
	return err
}

type chanClient struct {
	id    int
	clist chan []byte
}

func newChanClient(id int) *chanClient {
	c := &chanClient{
		id:    id,
		clist: make(chan []byte, 12),
	}
	go c.sendData()
	return c
}

func (c *chanClient) sendData() {
	var cnt int
	for {
		<-c.clist
		cnt++
		if cnt%200 == 0 && c.id == 1500 {
			fmt.Println(">>>>>> client", c.id, "dispatched ", cnt, "clist len ", len(c.clist))
		}
	}
}

func chanTest(cs int) {
	types.G.Go(func() error {
		chanStartPut()
		return nil
	})

	types.G.Go(func() error {
		chanDispatch()
		return nil
	})

	types.G.Go(func() error {
		chanAppendClient(cs)
		return nil
	})

	types.G.Wait()
}

func chanStartPut() {
	ticker := time.NewTicker(20 * time.Millisecond)
	var cnt int
	for {
		select {
		case <-ticker.C:
			select {
			case upchan <- []byte(types.DefaultString):
				cnt++
				if cnt%100 == 0 {
					fmt.Println(">>>> ", cnt)
				}
			default:
				fmt.Println("upchan overflow...")
			}
		}
	}
}

func chanDispatch() {
	for {
		select {
		case buf := <-upchan:
			clients.Each(func(v interface{}) bool {
				client, _ := v.(*chanClient)
				select {
				case client.clist <- buf:
				default:
					fmt.Println("client", client.id, "overflow...")
				}
				return false
			})
		}
	}
}

func chanAppendClient(cs int) {
	var cnt int
	for ; cnt < cs; cnt++ {
		c := newChanClient(cnt)
		clients.Add(c)
	}
}

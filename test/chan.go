package test

import (
	"fmt"
	"time"

	"github.com/scythefly/orb"
	"golang.org/x/sync/errgroup"
)

var (
	upchan  = make(chan []byte, 12)
	gChan   errgroup.Group
	clients = orb.NewSet()
)

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

func ChanTest(cs int) {
	gChan.Go(func() error {
		chanStartPut()
		return nil
	})

	gChan.Go(func() error {
		chanDispatch()
		return nil
	})

	gChan.Go(func() error {
		chanAppendClient(cs)
		return nil
	})

	gChan.Wait()
}

func chanStartPut() {
	ticker := time.NewTicker(20 * time.Millisecond)
	var cnt int
	for {
		select {
		case <-ticker.C:
			select {
			case upchan <- []byte(defaultString):
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

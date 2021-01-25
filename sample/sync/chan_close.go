package sync

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"auxx/types"
)

var (
	errSendTooMuch     = errors.New("send too much")
	errNoPacketInCache = errors.New("no packet in cache")
	errClose           = errors.New("server closed")
)

func newChanCloseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "close",
		Short: "Run chan close examples",
		RunE:  chanCloseRun,
	}

	return cmd
}

func chanCloseRun(_ *cobra.Command, _ []string) error {
	var err error
	xx := make(chan bool)
	types.G.Go(func() error {
		watch(xx)
		fmt.Println("----------")
		return nil
	})

	time.Sleep(2 * time.Second)
	fmt.Println(">>>>>>> publish!!")
	close(xx)

	types.G.Wait()
	return err
}

func watch(xx chan bool) {
	// first meta
	select {
	case v, ok := <-xx:
		if !ok {
			fmt.Println("xx closed >>", v)
		} else {
			fmt.Println("get from xx", v)
		}
	case <-time.After(5 * time.Second):
		fmt.Println("tttttttttimeout")
	}
	// read
	var cnt = 0
	var err error
	for {
		cnt++
		if err = readPacket(cnt); err != nil {
			fmt.Println(err)
			switch err {
			case errClose:
				return
			case errSendTooMuch:
				time.Sleep(2 * time.Second)
			case errNoPacketInCache:
				time.Sleep(40 * time.Millisecond)
			}
		}
	}
}

func readPacket(idx int) error {
	if idx%100 == 0 {
		return errClose
	} else if idx%30 == 0 {
		return errSendTooMuch
	} else if idx%10 == 0 {
		return errNoPacketInCache
	}
	return nil
}

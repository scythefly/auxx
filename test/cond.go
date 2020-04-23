package test

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	cond  *sync.Cond
	gCond errgroup.Group
)

func init() {
	cond = sync.NewCond(&sync.Mutex{})
}

func condWorker(x int) {
	cond.L.Lock()
	defer cond.L.Unlock()
	for i := 0; i < 5; i++ {
		cond.Wait()
		fmt.Println("worker", x, " gogogo!")
		time.Sleep(3 * time.Second)
	}
}

func CondTest() {
	for i := 0; i < 4; i++ {
		idx := i
		gCond.Go(func() error {
			condWorker(idx)
			return nil
		})
	}

	// time.Sleep(3 * time.Second)
	// cond.Signal()
	// time.Sleep(3 * time.Second)
	// cond.Signal()
	time.Sleep(3 * time.Second)
	fmt.Println("---- broadcast")
	cond.Broadcast()
	time.Sleep(1 * time.Second)
	fmt.Println("---- broadcast * 3")
	cond.Broadcast()
	cond.Broadcast()
	cond.Broadcast()
	time.Sleep(12 * time.Second)
	fmt.Println("---- broadcast")
	cond.Broadcast()
	gCond.Wait()
}

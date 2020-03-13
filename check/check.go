package check

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	tValue int
	once   sync.Once
	g      errgroup.Group
)

func Perform() {
	getValue()
}

func getTValue() int {
	once.Do(func() {
		time.Sleep(3 * time.Second)
		tValue = 10
	})
	return tValue
}

func getValue() {
	g.Go(func() error {
		fmt.Println(getTValue())
		return nil
	})

	g.Go(func() error {
		fmt.Println(getTValue())
		return nil
	})
	g.Wait()
}

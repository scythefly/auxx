package test

import (
	"context"
	"fmt"
	"time"
)

func CtxTest() {
	ctx := context.Background()
	cttx, cancel := context.WithCancel(ctx)
	for i := 0; i < 5; i++ {
		go wait(cttx)
	}
	time.Sleep(time.Second * 2)
	cancel()

	time.Sleep(time.Second)
}

func wait(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("wait goroutine quit...", ctx.Err())
	}
}

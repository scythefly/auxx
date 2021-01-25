package sync_test

import (
	"fmt"
	"testing"
	"time"
)

func Test_Chan(t *testing.T) {
	ch := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		closeChan(ch)
	}()
	go func() {
		for {
			time.Sleep(time.Second)
			sendToChan(ch)
		}
	}()
	for {
		select {
		case v, ok := <-ch:
			fmt.Println(">>>> case ok", ok)
			if ok {
				fmt.Println(">>> chan recv", v)
			} else {
				return
			}
		}
	}
}

func closeChan(ch chan int) {
	close(ch)
}

var idx int

func sendToChan(ch chan int) {
	ch <- idx
	idx++
}

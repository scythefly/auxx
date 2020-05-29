package test

import (
	"fmt"
	"time"
)

func TimerTest() {
	for {
		select {
		case <-time.After(time.Second):
			fmt.Printf("1s timer tick tick...\n")
		case <-time.After(5 * time.Second):
			fmt.Printf("5s timer tick tick tick...\n")
		}
	}
}

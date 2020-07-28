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

func TickerTest() {
	ticker1 := time.NewTicker(time.Second)
	ticker2 := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ticker1.C:
			fmt.Printf("1s timer tick tick...\n")
		case <-ticker2.C:
			fmt.Printf("5s timer tick tick...\n")
		}
	}
}

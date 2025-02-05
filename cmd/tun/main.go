// +build !windows

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"auxx/internal/tun"
)

func main() {
	dev, err := tun.CreateTUN("utun", 1420)
	if err == nil {
		realName, err2 := dev.Name()
		fmt.Println(realName, err2)
	} else {
		fmt.Println("createTUN", err)
	}

	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGTERM)
	signal.Notify(term, os.Interrupt)

	select {
	case <-term:
	}

	fmt.Println("===================== down =====================")
}

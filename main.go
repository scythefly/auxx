package main

import (
	"fmt"
	"runtime"

	"auxx/command"
)

func main() {
	runtime.GOMAXPROCS(12)
	if err := command.Execute(); err != nil {
		fmt.Println(err)
	}
}

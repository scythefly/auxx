package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	"auxx/command"
)

func main() {
	runtime.GOMAXPROCS(12)

	go http.ListenAndServe(":6060", nil)

	if err := command.Execute(); err != nil {
		fmt.Println(err)
	}
}

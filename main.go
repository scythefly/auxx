package main

import (
	"fmt"
	"math/rand"
	// _ "net/http/pprof"
	"runtime"
	"time"

	"auxx/command"
)

func main() {
	runtime.GOMAXPROCS(12)

	// go http.ListenAndServe(":6060", nil)

	rand.Seed(time.Now().UnixNano())
	if err := command.Execute(); err != nil {
		fmt.Println(err)
	}
}

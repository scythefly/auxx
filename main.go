package main

import (
	"runtime"

	"auxx/command"
)

func main() {
	runtime.GOMAXPROCS(12)
	command.Execute()
}

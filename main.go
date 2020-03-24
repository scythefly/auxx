package main

import (
	"runtime"

	"auxx/command"
)

func main() {
	runtime.GOMAXPROCS(4)
	command.Execute()
}

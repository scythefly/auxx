package main

import (
	"fmt"

	"grpc/hello"
)

func main() {
	fmt.Println(">>>>>>>>>>")
	var str hello.String
	fmt.Println(str.Value)
}

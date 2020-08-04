package test

import (
	"fmt"

	"github.com/baidu/go-lib/gotrack"
)

func PanicTest() {
	defer func() {
		fmt.Println(">>>>>> defer 1")
		if err := recover(); err != nil {
			fmt.Printf("panic: %s", gotrack.CurrentStackTrace(0))
		}
	}()

	defer func() {
		fmt.Println(">>>>>> defer 2")
		panic("panic in defer 2")
	}()

	panic("panic test")
}

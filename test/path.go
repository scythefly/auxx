package test

import (
	"fmt"
	"os"
	"path/filepath"
)

func PathTest() {
	p, err := filepath.Abs(os.Args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p)
	pp := filepath.Join(p, "../conf")
	fmt.Println(pp)
}

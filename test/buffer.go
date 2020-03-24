package test

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	gBuffer errgroup.Group
)

func Buffer() {
	buf := getBuffer(100)
	buf.Reset()
	buf.Write([]byte("12345"))
	gBuffer.Go(func() error {
		s := buf.Bytes()
		fmt.Println("get slice: ", string(s), len(s), buf.Len())
		time.Sleep(3 * time.Second)
		fmt.Println("get slice: ", string(s), len(s), buf.Len())
		return nil
	})

	gBuffer.Go(func() error {
		time.Sleep(time.Second)
		fmt.Println(">>> set 67890")
		buf.Reset()
		buf.Write([]byte("67890"))
		return nil
	})

	gBuffer.Wait()
	putBuffer(buf)
}

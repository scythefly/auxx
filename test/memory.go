package test

import (
	"bytes"
	"fmt"
	"time"
)

type pktNode struct {
	data       *bytes.Buffer
	prev, next *pktNode
}

func MemoryTest() {
	buf := make([]byte, 5*1024*1024)
	for idx, _ := range buf {
		buf[idx] = 0
	}
	node := &pktNode{
		data: bytes.NewBuffer(make([]byte, 5*1024*1024)),
	}
	node.data.Write(buf)
	n := node
	for i := 0; i < 20; i++ {
		node.next = &pktNode{
			data: bytes.NewBuffer(make([]byte, 5*1024*1024)),
			prev: node,
		}
		node = node.next
		node.data.Write(buf)
		if i == 50 {
			n = node
		}
	}
	time.Sleep(20 * time.Second)
	n.prev = nil
	fmt.Println("-----------------------------")
	stop := make(chan struct{})
	stop <- struct{}{}
}

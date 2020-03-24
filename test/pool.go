package test

import (
	"bytes"
	"sync"
)

var (
	pool1K = sync.Pool{
		New: func() interface{} { return &bytes.Buffer{} },
	}
	pool2K = sync.Pool{
		New: func() interface{} { return &bytes.Buffer{} },
	}
	pool4K = sync.Pool{
		New: func() interface{} { return &bytes.Buffer{} },
	}
)

func getBuffer(n int) *bytes.Buffer {
	if n < 1024 {
		return pool1K.Get().(*bytes.Buffer)
	} else if n < 2048 {
		return pool2K.Get().(*bytes.Buffer)
	} else if n < 4096 {
		return pool4K.Get().(*bytes.Buffer)
	}

	return nil
}

func putBuffer(b *bytes.Buffer) {
	if b.Cap() > 2048 {
		pool4K.Put(b)
	} else if b.Cap() > 1024 {
		pool2K.Put(b)
	} else {
		pool1K.Put(b)
	}
}

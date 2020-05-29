package test

import (
	"os"
	"sync"
	"testing"

	"golang.org/x/sync/errgroup"
)

var (
	gCondTest errgroup.Group
)

func Benchmark_CondLock1(b *testing.B) {
	cond := sync.NewCond(&sync.RWMutex{})
	f, err := os.Create("/dev/null")
	if err != nil {
		return
	}
	defer f.Close()
	for i := 0; i < b.N; i++ {
		gCondTest.Go(func() error {
			cond.L.Lock()
			f.Write([]byte(defaultString))
			f.Write([]byte(defaultString))
			f.Write([]byte(defaultString))
			cond.L.Unlock()
			return nil
		})
	}
	gCondTest.Wait()
}

func Benchmark_CondNoLock(b *testing.B) {
	// cond := sync.NewCond(&sync.RWMutex{})
	f, err := os.Create("/dev/null")
	if err != nil {
		return
	}
	for i := 0; i < b.N; i++ {
		gCondTest.Go(func() error {
			f.Write([]byte(defaultString))
			f.Write([]byte(defaultString))
			f.Write([]byte(defaultString))
			return nil
		})
	}
	gCondTest.Wait()
}

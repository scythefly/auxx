package test

import (
	"bytes"
	"testing"

	"golang.org/x/sync/errgroup"
)

var (
	g1 errgroup.Group
	g2 errgroup.Group
)

func Benchmark_UnusePool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		capt := i % 4096
		g1.Go(func() error {
			if capt < 100 {
				capt = 100
			}
			buf := &bytes.Buffer{}
			for buf.Len()+6 < capt {
				buf.Write([]byte("12345"))
			}
			return nil
		})
	}
	g1.Wait()
}

func Benchmark_UsePool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		capt := i % 4096
		g2.Go(func() error {
			if capt < 100 {
				capt = 100
			}
			buf := getBuffer(capt)
			buf.Reset()
			for buf.Len()+6 < capt {
				buf.Write([]byte("12345"))
			}
			putBuffer(buf)
			return nil
		})
	}
	g2.Wait()
}

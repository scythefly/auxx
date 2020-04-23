package test

import (
	"bytes"
	"context"
	"runtime"
	"testing"

	pool "github.com/jolestar/go-commons-pool/v2"
)

func Benchmark_NoPool(b *testing.B) {
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
	runtime.GC()
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
	runtime.GC()
}

func Benchmark_CommonPool(b *testing.B) {
	factory := pool.NewPooledObjectFactorySimple(
		func(context.Context) (interface{}, error) {
			return &bytes.Buffer{}, nil
		})
	ctx := context.Background()
	p := pool.NewObjectPoolWithDefaultConfig(ctx, factory)

	for i := 0; i < b.N; i++ {
		capt := i % 4096
		g3.Go(func() error {
			obj, err := p.BorrowObject(ctx)
			if err != nil {
				panic(err)
			}
			buf := obj.(*bytes.Buffer)
			buf.Reset()
			for buf.Len()+6 < capt {
				buf.Write([]byte("12345"))
			}
			if err = p.ReturnObject(ctx, obj); err != nil {
				panic(err)
			}
			return err
		})
	}
	g3.Wait()
	runtime.GC()
}

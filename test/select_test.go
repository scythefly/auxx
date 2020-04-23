package test

import (
	"context"
	"testing"
)

func Benchmark_Select(b *testing.B) {
	ctx, _ := context.WithCancel(context.Background())
	for i := 0; i < b.N; i++ {
		select {
		case <-ctx.Done():
		default:
		}
	}
}

func Benchmark_Select_if(b *testing.B) {
	var ok bool = true
	for i := 0; i < b.N; i++ {
		if ok {
		}
	}
}

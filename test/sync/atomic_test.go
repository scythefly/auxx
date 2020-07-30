package sync_test

import (
	"sync/atomic"
	"testing"
)

func Benchmark_Add(b *testing.B) {
	var v int64
	for i := 0; i < b.N; i++ {
		v += 100
	}
	b.Log(b.N, v)
}

func Benchmark_AtomicAdd(b *testing.B) {
	var v int64
	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&v, 100)
	}
	b.Log(b.N, v)
}

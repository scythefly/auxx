package test

import (
	"testing"
	"time"
)

func Benchmark_Time1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = time.Now()
	}
}

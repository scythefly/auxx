package test_test

import (
	"fmt"
	"testing"
)

func Benchmark_IdxSlice(b *testing.B) {
	var is []int
	for i := 0; i < 15; i++ {
		is = append(is, i)
	}

	for i := 0; i < b.N; i++ {
		_ = is[i%15]
	}
}

func Benchmark_IdxMap(b *testing.B) {
	is := make(map[string]int)
	for i := 0; i < 15; i++ {
		is[fmt.Sprintf("%d", i)] = i
	}

	for i := 0; i < b.N; i++ {
		_, _ = is["1"]
	}
}

// func Benchmark_Switch(b *testing.B) {
// 	is := make(map[int]int)
// 	for i := 0; i < 15; i++ {
// 		is[i] = i
// 	}
//
// 	for i := 0; i < b.N; i++ {
// 		switch i%15 {
// 		case 0
// 		}
// 	}
// }

package feature_test

import "testing"

func Benchmark_Closures(b *testing.B) {
	var a int64
	for i := 0; i < b.N; i++ {
		f := func() {
			a++
		}
		f()
	}
	b.Log(a)
}

func Benchmark_Closures1(b *testing.B) {
	var a int64
	f := func() {
		a++
	}
	for i := 0; i < b.N; i++ {
		f()
	}
	b.Log(a)
}

func add(a *int64) {
	*a++
}

func Benchmark_UnClosures(b *testing.B) {
	var a int64
	for i := 0; i < b.N; i++ {
		add(&a)
	}
	b.Log(a)
}

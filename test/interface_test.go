package test

import (
	"testing"
)

func withCommon(f *frame) bool {
	return f.idx < 100
}

func withInterface(idx interface{}) bool {
	p, _ := idx.(*frame)
	return p.idx < 100
}

func Benchmark_PackInterface(b *testing.B) {
	f := &frame{
		idx:   0,
		value: 1,
	}
	for i := 0; i < b.N; i++ {
		f.idx = i
		withInterface(f)
	}
}

func Benchmark_Unpack(b *testing.B) {
	f := &frame{
		idx:   0,
		value: 1,
	}
	for i := 0; i < b.N; i++ {
		f.idx = i
		withCommon(f)
	}
}

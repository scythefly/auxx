package sys_test

import "testing"

type sec1 struct {
	buf  [128]byte
	sync bool
}

type sec2 struct {
	buf  [4096]byte
	sync bool
}

type sec3 struct {
	buf  [128]int
	sync bool
}

type sec4 struct {
	buf  [4096]int
	sync bool
}

func Benchmark_AllocSec1(b *testing.B) {
	// var ss []*sec1
	for i := 0; i < b.N; i++ {
		s := &sec1{}
		s.sync = true
		// ss = append(ss, s)
	}
}

func Benchmark_AllocSec2(b *testing.B) {
	// var ss []*sec2
	for i := 0; i < b.N; i++ {
		s := &sec2{}
		s.sync = true
		// ss = append(ss, s)
	}
}

func Benchmark_AllocSec3(b *testing.B) {
	// var ss []*sec2
	for i := 0; i < b.N; i++ {
		s := &sec3{}
		s.sync = true
		// ss = append(ss, s)
	}
}

func Benchmark_AllocSec4(b *testing.B) {
	// var ss []*sec2
	for i := 0; i < b.N; i++ {
		s := &sec4{}
		s.sync = true
		// ss = append(ss, s)
	}
}
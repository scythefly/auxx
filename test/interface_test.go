package test

import (
	"testing"
)

type IA interface {
	AA()
}

type AA struct {
	idx int
}

func (a *AA) AA() {
	a.idx++
}

func interfaceRun(ia IA) {
	ia.AA()
}

func aaRun(a *AA) {
	a.AA()
}

func switchRun(v interface{}) {
	v.(IA).AA()
}

func Benchmark_InterfaceRun(b *testing.B) {
	a := &AA{}
	for i := 0; i < b.N; i++ {
		interfaceRun(a)
	}
}

func Benchmark_Run(b *testing.B) {
	a := &AA{}
	for i := 0; i < b.N; i++ {
		aaRun(a)
	}
}

func Benchmark_switchRun(b *testing.B) {
	a := &AA{}
	for i := 0; i < b.N; i++ {
		switchRun(a)
	}
}

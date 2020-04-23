package test

import (
	"bufio"
	"os"
	"testing"
)

func Benchmark_ReadBytes_bufio_1024(b *testing.B) {
	f, err := os.Open("/Users/iuz/Downloads/download/fcc.txt")
	if err != nil {
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)
	buf := make([]byte, 1024)
	for i := 0; i < b.N; i++ {
		r.Read(buf)
	}
}

func Benchmark_ReadBytes_bufio_4096(b *testing.B) {
	f, err := os.Open("/Users/iuz/Downloads/download/fcc.txt")
	if err != nil {
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)
	buf := make([]byte, 4096)
	for i := 0; i < b.N; i++ {
		r.Read(buf)
	}
}

func Benchmark_ReadBytes_file_4096(b *testing.B) {
	f, err := os.Open("/Users/iuz/Downloads/download/fcc.txt")
	if err != nil {
		return
	}
	defer f.Close()
	buf := make([]byte, 4096)
	for i := 0; i < b.N; i++ {
		f.Read(buf)
	}
}

func Benchmark_ReadBytes_file_1024(b *testing.B) {
	f, err := os.Open("/Users/iuz/Downloads/download/fcc.txt")
	if err != nil {
		return
	}
	defer f.Close()
	buf := make([]byte, 1024)
	for i := 0; i < b.N; i++ {
		f.Read(buf)
	}
}

func Benchmark_Write_1(b *testing.B) {
	f, err := os.Create("/dev/null")
	if err != nil {
		return
	}
	defer f.Close()

	for i := 0; i < b.N; i++ {
		f.Write([]byte(defaultString))
	}
}

func Benchmark_Write_2(b *testing.B) {
	f, err := os.Create("/dev/null")
	if err != nil {
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	for i := 0; i < b.N; i++ {
		w.Write([]byte(defaultString))
	}
}
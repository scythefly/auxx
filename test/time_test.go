package test

import (
	"fmt"
	"testing"
	"time"
)

func Benchmark_Time1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = time.Now()
	}
}

func Test_TimeAPI(tt *testing.T) {
	var t time.Time

	fmt.Println(t.IsZero())
	fmt.Println(t.Unix())
	fmt.Println(t.Year(), t.Month(), t.Day())

	t = time.Now()
	fmt.Println(t.IsZero())
	fmt.Println(t.Unix())
	fmt.Println(t.Year(), t.Month(), t.Day())
}

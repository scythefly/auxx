package test

import "fmt"

func DeferTest() {
	fmt.Printf(">>> output: %d\n", deferReturn(12))
	f := deferSwap()
	fmt.Printf(">>> defer output: %d\n", f.idx)
}

func deferReturn(input int) (output int) {
	defer func() {
		output = input
	}()
	return 2
}

type deferTestClass struct {
	idx int
}

func deferSwap() (ff *deferTestClass) {
	var f *deferTestClass
	defer func() {
		ff = fixClass(f)
	}()

	f = &deferTestClass{idx: 10}
	return f
}

func fixClass(f *deferTestClass) *deferTestClass {
	if f == nil {
		return &deferTestClass{idx: 1}
	}

	return &deferTestClass{
		idx: f.idx + 1,
	}
}

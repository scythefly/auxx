package test

import "fmt"

type arrayTest struct {
	inBytes [60]uint64
}

type arrA [60]byte

const (
	arrayTestString = "FLV\x00\x00"
)

func ArrayTest() {
	in := &arrayTest{}

	for i := 0; i < 200; i++ {
		in.inBytes[i%60] = uint64(i)
	}

	var out uint64
	for _, v := range in.inBytes {
		out += v
	}

	fmt.Println(out)

	fmt.Println(arrayTestString)

	var aa arrA
	aa[0] = 10
	fmt.Println(aa[0])

	arrayPara(&aa)
	fmt.Println(aa[0])
}

func arrayPara(aa *arrA) {
	aa[0] = 20
}

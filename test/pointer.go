// Is assigning a pointer atomic in golang?
// Is assignint a int64 atomic in golang?
package test

import "fmt"

const MAX_INT64 = int64((^uint64(0)) >> 1)

var (
	vInt64      int64
	pInt64      *int64
	valueLoop   bool
	pointerLoop bool
)

func writeValue(value int64) {
	for valueLoop {
		vInt64 = value
	}
}

func swapPointer(p *int64) {
	for pointerLoop {
		pInt64 = p
	}
}

func PointerTest() {
	// assign int64
	vInt64 = 1
	valueLoop = true
	var cnt1, cnt2 int
	var pp3 int64
	go writeValue(1)
	go writeValue(MAX_INT64)

	for i := 0; i < 100000; i++ {
		pp3 = vInt64
		if pp3 == 1 {
			cnt1++
			continue
		}
		if pp3 == MAX_INT64 {
			cnt2++
			continue
		}
		fmt.Printf(">>> !!!!! assign value it not atomic !!!!!")
	}
	fmt.Printf(">>> %d - %d\n", cnt1, cnt2)
	valueLoop = false
	cnt1 = 0
	cnt2 = 0

	// assign pointer
	var pp1 int64 = 1
	var pp2 = MAX_INT64
	pointerLoop = true
	pInt64 = &pp1

	go swapPointer(&pp1)
	go swapPointer(&pp2)

	for i := 0; i < 100000; i++ {
		pp3 = *pInt64
		if pp3 == 1 {
			cnt1++
			continue
		}
		if pp3 == MAX_INT64 {
			cnt2++
			continue
		}
		fmt.Printf(">>> !!!!! assign pointer it not atomic !!!!!")
	}
	pointerLoop = false
	fmt.Printf(">>> %d - %d\n", cnt1, cnt2)
}

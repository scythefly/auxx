package test

import "testing"

func Test_MapGetValue(t *testing.T) {
	var vm = map[uint64]string{
		1: "1111",
		2: "2222",
	}

	s, _ := vm[1]
	t.Log(s)
	s, _ = vm[3]
	t.Log(s)
}

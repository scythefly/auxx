package test

import (
	"fmt"
	"reflect"
	"time"
)

var interfaceMap = make(map[string]interface{})

func InterfaceTest() {
	interfaceMap["123"] = 123
	var i1 int
	var i2 uint32
	var i3 []byte
	interfaceTypeOf(i1)
	interfaceTypeOf(i2)
	interfaceTypeOf(i3)

	interfaceModify(&i1)
	fmt.Println(i1)

	var ii1 interface{}
	ii1 = 1
	v, ok := ii1.(*int)
	fmt.Println(v, ok)

	iii := &interfaceSession{
		Name:     "iiiTest",
		App:      "iiiApp",
		t1:       time.Now(),
		t2:       time.Now().Add(time.Hour * 24),
		inBytes:  1010101010,
		outBytes: 202020202,
	}
	iii.init()
	fmt.Println(iii.GetVariables("Name"))
	fmt.Println(iii.GetVariables("App"))
	fmt.Println(iii.GetVariables("t1"))
	fmt.Println(iii.GetVariables("t2"))
	fmt.Println(iii.GetVariables("inBytes"))
	fmt.Println(iii.GetVariables("outBytes"))
}

func interfaceTypeOf(dst interface{}) {
	fmt.Println(reflect.TypeOf(dst))
}

func interfaceModify(dst interface{}) {
	switch dst.(type) {
	case *int:
		v := dst.(*int)
		*v = 2
	}
}

type interfaceSession struct {
	Name     string
	App      string
	t1       time.Time
	t2       time.Time
	inBytes  int64
	outBytes int64

	vmap map[string]interface{}
}

func (s *interfaceSession) init() {
	s.vmap = make(map[string]interface{})
	s.vmap["Name"] = &s.Name
	s.vmap["App"] = &s.App
	s.vmap["t1"] = &s.t1
	s.vmap["t2"] = &s.t2
	s.vmap["inBytes"] = &s.inBytes
	s.vmap["outBytes"] = &s.outBytes
}

func (s *interfaceSession) GetVariables(vn string) string {
	if v, ok := s.vmap[vn]; ok {
		switch v.(type) {
		case *string:
			return *v.(*string)
		case *time.Time:
			return v.(*time.Time).Format(time.RFC3339)
		case *int64:
			return fmt.Sprintf("%d", *v.(*int64))
		}
	}
	return "nil"
}

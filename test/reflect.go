package test

import (
	"fmt"
	"reflect"
	"time"
)

type reflectSession struct {
	Name     string
	App      string
	t1       time.Time
	t2       time.Time
	inBytes  int64
	outBytes int64
}

func ReflectTest() {
	rs := &reflectSession{
		Name:     "iiiTest",
		App:      "iiiApp",
		t1:       time.Now(),
		t2:       time.Now().Add(time.Hour * 24),
		inBytes:  1010101010,
		outBytes: 202020202,
	}

	fmt.Println(rs.GetVariables("Name"))
	fmt.Println(rs.GetVariables("App"))
	fmt.Println(rs.GetVariables("t1"))
	fmt.Println(rs.GetVariables("t2"))
	fmt.Println(rs.GetVariables("inBytes"))
	fmt.Println(rs.GetVariables("outBytes"))
}

func (r *reflectSession) GetVariables(vn string) string {
	sv := reflect.ValueOf(r)
	v := sv.FieldByName(vn).Interface()
	switch v.(type) {
	case string:
		return v.(string)
	case time.Time:
		return v.(time.Time).Format(time.RFC3339)
	case int64:
		return fmt.Sprintf("%d", v.(int64))
	}
	return "nil"
}

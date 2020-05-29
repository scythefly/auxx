package test

import (
	"fmt"
	"net/url"
	"strconv"
)

func URLTest() {
	urlString := "test?serverId=scythefly.top"
	u, err := url.Parse(urlString)
	if err != nil {
		fmt.Printf("parse return err: %s\n", err.Error())
	}
	fmt.Println(u.Path, u.RawQuery)
	fmt.Println(u.Query())

	v, _ := strconv.Atoi(u.Query().Get("genId"))
	fmt.Println(v)
}

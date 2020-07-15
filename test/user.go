package test

import (
	"fmt"
	"os/user"
)

func UserTest() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Uid: %s, Gid: %s, Username: %s, Name: %s, HomeDir: %s\n",
		usr.Uid, usr.Gid, usr.Username, usr.Name, usr.HomeDir)
}

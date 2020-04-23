package test

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	user2 "os/user"
	"regexp"
)

func XmlTest() {
	reg := regexp.MustCompile(`^\$USER_HOME\$(.*)$`)
	var userHome string
	user, err := user2.Current()
	if err != nil {
		fmt.Println(err)
	} else {
		userHome = user.HomeDir
	}
	fmt.Println(user)
	if userHome == "" {
		userHome = "/Users/iuz"
	}
	conf, err := os.Open(userHome + "/Library/ApplicationSupport/JetBrains/GoLand2020.1/options/recentProjects.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conf.Close()
	data, err := ioutil.ReadAll(conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	app := application{}
	err = xml.Unmarshal(data, &app)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(app.Component.Option)
	for _, opt := range app.Component.Option {
		if opt.Name == "recentPaths" {
			for _, l := range opt.List.Value {
				fmt.Println(reg.ReplaceAllString(l.Value, "/Users/iuz${1}"))
			}
		}
	}
}

type Option struct {
	Name string `xml:"name,attr"`
	List List   `xml:"list"`
}

type application struct {
	XMLName   xml.Name  `xml:"application"`
	Component Component `xml:"component"`
}

type Component struct {
	Name   string   `xml:"name,attr"`
	Option []Option `xml:"option"`
}

type List struct {
	Value []Value `xml:"option"`
}

type Value struct {
	Value string `xml:"value,attr"`
}

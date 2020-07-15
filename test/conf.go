package test

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/pkg/errors"
)

func ConfTest() {
	setDefaultConfig()

	testConf("./test/data/mod_sample_conf.json")
	testConf("./test/data/mod_sample_conf_1.json")
}

func testConf(path string) {
	cc, err := loadConf(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	conf := cc.getConf("test1.scythefly.top", "app1/stream1")
	fmt.Println(conf.Desc)
}

type Config struct {
	Desc string `json:"desc"`
}

var defaultConf Config
var modulesConfPath string

type Location struct {
	Path   string  `json:"path"`
	Config *Config `json:"config"`

	reg *regexp.Regexp
}
type Hosts map[string][]*Location
type ModConf struct {
	Version  string
	ServerID Hosts
}

func (cc *ModConf) getConf(serverID, path string) *Config {
	var conf *Config
	if lts, ok := cc.ServerID[serverID]; ok {
		for _, lt := range lts {
			if lt.reg.Match([]byte(path)) {
				conf = lt.Config
				break
			}
		}
	}
	if conf == nil {
		conf = &defaultConf
	}

	return conf
}

func setDefaultConfig() {
	defaultConf.Desc = "This is default config"
}

func loadConf(path string) (*ModConf, error) {
	var conf ModConf
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&conf); err != nil {
		return nil, err
	}
	if conf.ServerID == nil {
		return nil, errors.New("no valid server id config")
	}

	for sid, lts := range conf.ServerID {
		for _, lt := range lts {
			(*lt).reg, err = regexp.Compile((*lt).Path)
			if err != nil {
				return nil, errors.WithMessagef(err, "[%s] parse regexp '%s'", sid, lt.Path)
			}
		}
	}

	return &conf, nil
}

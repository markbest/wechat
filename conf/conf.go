package conf

import (
	"errors"
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

var (
	Conf       config
	ConfigFile = "./conf/conf.toml"
)

type config struct {
	App app `toml:"app"`
}

type app struct {
	Port  string
	Token string
}

func InitConfig() error {
	configBytes, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		return errors.New("config load err:" + err.Error())
	}
	_, err = toml.Decode(string(configBytes), &Conf)
	if err != nil {
		return errors.New("config decode err:" + err.Error())
	}
	return nil
}

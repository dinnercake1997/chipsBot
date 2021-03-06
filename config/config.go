package config


import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var Myconfig Config

type Config struct {
	MiraiHttpUrl      string`yaml:"MiraiHttpUrl"`
	Auth   string  `yaml:"Auth"`
	SessionKey string `yaml:"SessionKey"`
	QQNumber string `yaml:"QQNumber"`
}

func (c *Config)GetConf() *Config {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}




package config


import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var Myconfig Config
var QQSendGroups=[]string{"1085171553","763091038"}
var SendOnTitleMap map[string]string

type Config struct {
	MiraiHttpUrl      string`yaml:"MiraiHttpUrl"`
	Auth   string  `yaml:"Auth"`
	SessionKey string `yaml:"SessionKey"`
	QQNumber string `yaml:"QQNumber"`
	VersionTime string `yaml:"VersionTime"`
}

func (c *Config)GetConf() *Config {
	SendOnTitleMap= make(map[string]string)
	for _, v := range QQSendGroups {
		SendOnTitleMap[v]=""
	}
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




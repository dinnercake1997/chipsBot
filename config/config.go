package config


import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)


var Myconfig Config
var QQSendGroups=[]string{"1085171553","763091038"}
var GroupSets =make(map[string] *GroupSet)

type GroupSet struct{
	Tittle string
	IsOnTime bool
}

type Config struct {
	MiraiHttpUrl      string`yaml:"MiraiHttpUrl"`
	Auth   string  `yaml:"Auth"`
	SessionKey string `yaml:"SessionKey"`
	QQNumber string `yaml:"QQNumber"`
	VersionTime string `yaml:"VersionTime"`
}


func (c *Config)GetConf() *Config {
	GroupSets= make(map[string]*GroupSet)
	for _, v := range QQSendGroups {
		tempSet:=new(GroupSet)
		tempSet.Tittle="随机"
		tempSet.IsOnTime=false
		GroupSets[v]=tempSet
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




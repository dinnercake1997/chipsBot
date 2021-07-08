package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)


var Myconfig Config

var GroupSets =make(map[string] *GroupSet)

type GroupSet struct{
	Tittle string
	IsOnTime bool
	IsWeiBoFuLiJi bool
	IsWeiBoSeTu bool
	IsWeiBoShaDiaoTU bool
}

type Config struct {
	MiraiHttpUrl      string`yaml:"MiraiHttpUrl"`
	Auth   string  `yaml:"Auth"`
	SessionKey string `yaml:"SessionKey"`
	QQNumber string `yaml:"QQNumber"`
	VersionTime string `yaml:"VersionTime"`
	TargetGroups []string `yaml:"TargetGroups"`
	IsNewWeiBoSeconds int64 `yaml:"IsNewWeiBoSeconds"`
	WeiBoShaDiaoUps []string `yaml:"WeiBoShaDiaoUps"`
	WeiBoFuLiJiUps []string `yaml:"WeiBoFuLiJiUps"`
	WeiBoSeTuUps []string `yaml:"WeiBoSeTuUps"`
}


func (c *Config)GetConf() *Config {
	GroupSets= make(map[string]*GroupSet)


	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, v := range Myconfig.TargetGroups  {
		tempSet:=new(GroupSet)
		tempSet.Tittle="随机"
		tempSet.IsOnTime=false
		GroupSets[v]=tempSet
	}
	return c
}




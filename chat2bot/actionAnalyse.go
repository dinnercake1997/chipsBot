package chat2bot

import (
	"chipsBot/BotService"
	"log"
	"regexp"
)
func ActionSelect(text string)(err error)  {

	reg:=regexp.MustCompile("[来给求][点张份](?P<keyword>.*?)?[色涩瑟][图]")
	key:=reg.FindSubmatch([]byte(text))
	log.Printf("ActionSelectKey:%v\n",string(key[1]))
	if len(key[0])==0{
		//没有识别出
		return
	}
	if len(key[1])==0{
		BotService.SendPic()
	}else{
		BotService.SendPicWithKey(string(key[1]))
	}
	return
}
package chat2bot

import (
	"errors"
	"fmt"
	//"github.com/NateScarlet/pixiv/pkg/novel"
	"github.com/bitly/go-simplejson"
	"log"
)

func GetTextFromBody(bodyString string)(text string ,err error){
	json,err:= simplejson.NewJson([]byte(bodyString))
	if err!=nil{
		return
	}
	//body:{"type":"GroupMessage",
	//"messageChain":[{"type":"Source","id":237890,"time":1614921415},{"type":"Plain","text":".。。"}],
	//"sender":{"id":1020224260,"memberName":"色之巨人-薯",
	//"permission":"MEMBER",
	//"group":{"id":763091038,
	//"name":"吾嘴与汝垢，孰臭？",
	//"permission":"MEMBER"}}}
	messageChainArray, err :=json.Get("messageChain").Array()
	for _,item :=range messageChainArray {
		dataMap,ok:= item.(map[string]interface{})
		if !ok{
			log.Printf("请求出错:%v",err)
			return "",errors.New("请求出错")
		}
		messageType,_:=dataMap["type"].(string)
		if messageType=="Plain"{
			text,_:=dataMap["text"].(string)
			return text,nil
		}
	}
	return text,err
}
func GetTextAndGroupFromBody(bodyString string)(text string ,groupQQ string ,err error){
	json,err:= simplejson.NewJson([]byte(bodyString))
	if err!=nil{
		return
	}
	//body:{"type":"GroupMessage",
	//"messageChain":[{"type":"Source","id":237890,"time":1614921415},{"type":"Plain","text":".。。"}],
	//"sender":{"id":1020224260,"memberName":"色之巨人-薯",
	//"permission":"MEMBER",
	//"group":{"id":763091038,
	//"name":"吾嘴与汝垢，孰臭？",
	//"permission":"MEMBER"}}}
	messageChainArray, err :=json.Get("messageChain").Array()
	groupqq,_:=json.Get("sender").Get("group").Map()
	//groupQQ=groupQQInt.(string)
	groupQQ=fmt.Sprintf("%v",groupqq["id"])
	log.Printf("groupQQ为:%s",groupQQ)
	//group:=sender.Get("group")
	//groupQQ, _ =group.Get("id").String()
	for _,item :=range messageChainArray {
		dataMap,ok:= item.(map[string]interface{})
		if !ok{
			log.Printf("请求出错:%v",err)
			return "","",errors.New("请求出错")
		}
		messageType,_:=dataMap["type"].(string)
		if messageType=="Plain"{
			text,_:=dataMap["text"].(string)

			return text,groupQQ,nil
		}
	}

	return text,groupQQ,err
}


//body:{"type":"GroupMessage",
//"messageChain":[{"type":"Source","id":115266,"time":1624519250},{"type":"Plain","text":"test"}],
//"sender":{"id":1020224260,
//"memberName":"薯道之难难于上青天","permission":"MEMBER","group":{"id":1085171553,"name":"BCR-昏睡的聆时之森","permission":"MEMBER"}}}


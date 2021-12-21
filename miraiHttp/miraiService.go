package miraiHttp

import (
	"chipsBot/config"
	"chipsBot/utils"
	"errors"
	"github.com/bitly/go-simplejson"
	"log"
	"time"
)
var SessionKey string

func authMiraiHttp()(session string ,err error){
	params:="{\n    \"authKey\": \""+config.Myconfig.Auth+"\"}"
	content,err:=utils.DoPost(config.Myconfig.MiraiHttpUrl+"auth",params)
	if err!=nil{
		log.Printf("请求初始化mirai-http-api出错:%v",err)
		return "",err
	}

	log.Printf("content:%v",content)
	json,err:= simplejson.NewJson([]byte(content))
	if err!=nil{
		log.Printf("请求:%v",err)
		return "",errors.New("请求初始化mirai-http-api转json出错")
	}
	if err!=nil{return}
	session, err =json.Get("session").String()
	if err!=nil{
		log.Printf("请求初始化mirai-http-api获取session:%v",err)
		return "",errors.New("请求初始化mirai-http-api获取session出错")
	}
	return session,nil
}

func verifyMiraiHttp(session string)(err error){
	params:="{\n    \"sessionKey\": \""+session+"\",\n    \"qq\" : "+config.Myconfig.QQNumber+"}"
	_,err=utils.DoPost(config.Myconfig.MiraiHttpUrl+"verify",params)
	if err!=nil{
		log.Printf("请求初始化mirai-http-api出错:%v",err)
		return
	}
	SessionKey=session
	return nil
}


func InitMiraiHttp()error{

	session,err:=authMiraiHttp()
	if err!=nil{
		log.Printf("初始化miraihttp出错")
		return err
	}
	verifyMiraiHttp(session)
	return nil
}
func SendPic( picUrl string ){

	params:="{\n    \"sessionKey\": \""+SessionKey+"\",\n    \"target\": 1085171553,\n    \"messageChain\": [\n        { \"type\": \"Image\", \"url\": \""+picUrl+"\" }\n    ]\n}"
	_,err:=utils.DoPost(config.Myconfig.MiraiHttpUrl+"sendGroupMessage",params)
	//log.Printf("请求发送图片参数为：%v",params)
	if err!=nil{
		log.Printf("请求mirai-http-api出错:%v",err)
		return
	}
}

func SendPicByQQ( picUrl string ,qq string){

	params:="{\n    \"sessionKey\": \""+SessionKey+"\",\n    \"target\": "+qq+",\n   \"messageChain\": [\n        { \"type\": \"Image\", \"url\": \""+picUrl+"\" }\n    ]\n}"
	_,err:=utils.DoPost(config.Myconfig.MiraiHttpUrl+"sendGroupMessage",params)
	//log.Printf("请求发送图片参数为：%v",params)
	if err!=nil{
		log.Printf("请求mirai-http-api出错:%v",err)
		return
	}
}
func SendPicByQQWithCheHui( picUrl string ,qq string)(err error){

	params:="{\n    \"sessionKey\": \""+SessionKey+"\",\n    \"target\": "+qq+",\n   \"messageChain\": [\n        { \"type\": \"Image\", \"url\": \""+picUrl+"\" }\n    ]\n}"
	content,err:=utils.DoPost(config.Myconfig.MiraiHttpUrl+"sendGroupMessage",params)
	//log.Printf("请求发送图片参数为：%v",params)
	if err!=nil{
		log.Printf("请求mirai-http-api出错:%v",err)
		return
	}
	res:=[]byte(content)
	resJson,err:=simplejson.NewJson(res)
	msgId,err:=resJson.Get("messageId").String()

	//10秒后撤回
	time.Sleep(10 * time.Second)

	params="{\n    \"sessionKey\": \""+SessionKey+"\",\n    \"target\": "+msgId+"\n}"
	content,err=utils.DoPost(config.Myconfig.MiraiHttpUrl+"/recall",params)
	log.Printf("撤回结果：%v",content)
	if err!=nil{
		log.Printf("撤回出错:%v",err)
		return
	}
	return
}


func SendText( text string ){
	params:="{\n    \"sessionKey\": \""+SessionKey+"\",\n   \"target\": 1085171553,\n    \"messageChain\": [\n        { \"type\": \"Plain\"," +
		" \"text\": \""+text+"\" }\n    ]\n}"
	_,err:=utils.DoPost(config.Myconfig.MiraiHttpUrl+"sendGroupMessage",params)
	if err!=nil{
		log.Printf("请求mirai-http-api出错:%v",err)
		return
	}
	return
}


func SendTextByQQ( text string ,qq string){
	params:="{\n    \"sessionKey\": \""+SessionKey+"\",\n    \"target\": "+qq+",\n    \"messageChain\": [\n        { \"type\": \"Plain\"," +
		" \"text\": \""+text+"\" }\n    ]\n}"
	_,err:=utils.DoPost(config.Myconfig.MiraiHttpUrl+"sendGroupMessage",params)
	if err!=nil{
		log.Printf("请求mirai-http-api出错:%v",err)
		return
	}
	return
}

func SendMix( text string ,url string ){

	params:="{\n    \"sessionKey\": \""+SessionKey+"\",\n    \"target\": 1085171553,\n    \"messageChain\": [\n                { \"type\": \"Plain\", \"text\": \""+text+"\" }\n    ]\n}"
	//params2:="{\n    \"sessionKey\": \""+SessionKey+"\",\n    \"target\": 763091038,\n    \"messageChain\": [\n        { \"type\": \"Image\", \"url\": \""+url+"\" },\n         { \"type\": \"Plain\", \"text\": \""+text+"\" }\n    ]\n}"
	_,err:=utils.DoPost(config.Myconfig.MiraiHttpUrl+"sendGroupMessage",params)
	if err!=nil{
		log.Printf("请求mirai-http-api出错:%v",err)
		return
	}
	return
}
func SendMixByQQ( text string ,qq string ){

	params:="{\n    \"sessionKey\": \""+SessionKey+"\",\n    \"target\": "+qq+",\n   \"messageChain\": [\n                { \"type\": \"Plain\", \"text\": \""+text+"\" }\n    ]\n}"
	//params2:="{\n    \"sessionKey\": \""+SessionKey+"\",\n    \"target\": 763091038,\n    \"messageChain\": [\n        { \"type\": \"Image\", \"url\": \""+url+"\" },\n         { \"type\": \"Plain\", \"text\": \""+text+"\" }\n    ]\n}"
	_,err:=utils.DoPost(config.Myconfig.MiraiHttpUrl+"sendGroupMessage",params)
	if err!=nil{
		log.Printf("请求mirai-http-api出错:%v",err)
		return
	}
	return
}
func SendPicAndTextByQQ( text string ,qq string ,url string){

	//params:="{\n    \"sessionKey\": \""+SessionKey+"\",\n    \"target\": "+qq+",\n   \"messageChain\": [\n                { \"type\": \"Plain\", \"text\": \""+text+"\" }\n    ]\n}"
	params2:="{\n    \"sessionKey\": \""+SessionKey+"\",\n    \"target\": "+qq+",\n    \"messageChain\": [\n        { \"type\": \"Image\", \"url\": \""+url+"\" },\n         { \"type\": \"Plain\", \"text\": \""+text+"\" }\n    ]\n}"
	_,err:=utils.DoPost(config.Myconfig.MiraiHttpUrl+"sendGroupMessage",params2)
	if err!=nil{
		log.Printf("请求mirai-http-api出错:%v",err)
		return
	}
	return
}
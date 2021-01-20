package miraiHttp

import (
	"chipsBot/utils"
	"log"
)

func SendPic( picUrl string ){

	params:="{\n    \"sessionKey\": \"q3j9t3SM\",\n    \"target\": 763091038,\n    \"messageChain\": [\n        { \"type\": \"Image\", \"url\": \""+picUrl+"\" }\n    ]\n}"
	_,err:=utils.DoPost("http://49.235.237.247:8080/sendGroupMessage",params)
	if err!=nil{
		log.Printf("请求mirai-http-api出错:%v",err)
		return
	}
}

func SendText( text string ){
	params:="{\n    \"sessionKey\": \"q3j9t3SM\",\n    \"target\": 763091038,\n    \"messageChain\": [\n        { \"type\": \"Plain\"," +
		" \"text\": \""+text+"\" }\n    ]\n}"
	_,err:=utils.DoPost("http://49.235.237.247:8080/sendGroupMessage",params)
	if err!=nil{
		log.Printf("请求mirai-http-api出错:%v",err)
		return
	}
	return
}

func SendMix( text string ,url string ){

	params:="{\n    \"sessionKey\": \"q3j9t3SM\",\n    \"target\": 763091038,\n    \"messageChain\": [\n        { \"type\": \"Image\", \"url\": \""+url+"\" },\n         { \"type\": \"Plain\", \"text\": \""+text+"\" }\n    ]\n}"
	_,err:=utils.DoPost("http://49.235.237.247:8080/sendGroupMessage",params)
	if err!=nil{
		log.Printf("请求mirai-http-api出错:%v",err)
		return
	}
	return
}
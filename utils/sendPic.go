package utils

import "log"

func SendPic(){
	//picUrl:="https://i.pixiv.cat/img-original/img/2014/04/05/23/26/27/42702642_p0.jpg"
	picUrl,err:=GetPicUrl()
	if err!=nil{
		log.Printf("GetPicUrl函数出错:%v",err)
		params:="{\n    \"sessionKey\": \"q3j9t3SM\",\n    \"target\": 763091038,\n    \"messageChain\": [\n        { \"type\": \"Plain\"," +
			" \"text\": \""+err.Error()+"\" }\n    ]\n}"
		_,err=DoPost("http://49.235.237.247:8080/sendGroupMessage",params)
		return
	}
	params:="{\n    \"sessionKey\": \"q3j9t3SM\",\n    \"target\": 763091038,\n    \"messageChain\": [\n        { \"type\": \"Image\", \"url\": \""+picUrl+"\" }\n    ]\n}"
	_,err=DoPost("http://49.235.237.247:8080/sendGroupMessage",params)
	if err!=nil{
		log.Printf("请求mirai-http-api出错:%v",err)
		return
	}

}

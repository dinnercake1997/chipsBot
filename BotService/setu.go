package BotService

import (
	"chipsBot/miraiHttp"
	"chipsBot/utils"
	"errors"
	"github.com/bitly/go-simplejson"
	"log"
)

var selector struct{

}

//从luolikonAPI获取图片Url
func GetPicUrl()(picUrl string,err error){
	url:="https://api.lolicon.app/setu/?r18=1&size1200=true&apikey=908789905fb629c064ccf8"
	content,err:=utils.DoGet(url)
	if err!=nil{
		log.Printf("请求loliconAPIkey出错:%v",err)
		return "",errors.New("请求loliconAPIkey出错")
	}
	log.Printf("content:%v",content)
	json,err:= simplejson.NewJson([]byte(content))
	if err!=nil{
		log.Printf("请求loliconAPIkey转json出错:%v",err)
		return "",errors.New("请求loliconAPIkey转json出错")
	}
	if err!=nil{return}
	dataArrary, err :=json.Get("data").Array()
	if err!=nil{
		log.Printf("请求loliconAPIkey获取图片信息出错:%v",err)
		return "",errors.New("请求loliconAPIkey获取图片列表出错")
	}
	dataMap,ok:= dataArrary[0].(map[string]interface{})
	if !ok{
		log.Printf("请求loliconAPIkey获取图片url出错:%v",err)
		return "",errors.New("请求loliconAPIkey获取图片信息出错")
	}
	picUrl,ok=dataMap["url"].(string)
	if !ok{
		log.Printf("请求loliconAPIkey获取图片url出错:%v",err)
		return "",errors.New("请求loliconAPIkey获取图片url出错")
	}
	log.Printf("picUrl:%v",picUrl)
	return
	//dataArraryBytes,_:=dataArrary.Bytes()
	//dataArraryJson,_:= simplejson.NewJson(dataArraryBytes)
	//log.Printf("dataArraryJson:%s",dataArraryJson)

}


func SendPic(){
	//picUrl:="https://i.pixiv.cat/img-original/img/2014/04/05/23/26/27/42702642_p0.jpg"
	picUrl,err:=GetPicUrl()
	if err!=nil{
		log.Printf("GetPicUrl函数出错:%v",err)
		miraiHttp.SendText(err.Error())
		return
	}
	miraiHttp.SendPic(picUrl)
}

//从luolikonAPI获取图片Url
func GetPicUrlWithKey(key string)(picUrl string,err error){
	url:="https://api.lolicon.app/setu/?r18=1&size1200=true&apikey=908789905fb629c064ccf8"
	keyWord:="&keyword="+key
	url=url+keyWord
	content,err:=utils.DoGet(url)
	if err!=nil{
		log.Printf("请求loliconAPIkey出错:%v",err)
		return "",errors.New("请求loliconAPIkey出错")
	}
	log.Printf("content:%v",content)
	json,err:= simplejson.NewJson([]byte(content))
	if err!=nil{
		log.Printf("请求loliconAPIkey转json出错:%v",err)
		return "",errors.New("请求loliconAPIkey转json出错")
	}
	if err!=nil{return}
	dataArrary, err :=json.Get("data").Array()
	if err!=nil{
		log.Printf("请求loliconAPIkey获取图片信息出错:%v",err)
		return "",errors.New("请求loliconAPIkey获取图片列表出错")
	}
	dataMap,ok:= dataArrary[0].(map[string]interface{})
	if !ok{
		log.Printf("请求loliconAPIkey获取图片url出错:%v",err)
		return "",errors.New("请求loliconAPIkey获取图片信息出错")
	}
	picUrl,ok=dataMap["url"].(string)
	if !ok{
		log.Printf("请求loliconAPIkey获取图片url出错:%v",err)
		return "",errors.New("请求loliconAPIkey获取图片url出错")
	}
	log.Printf("picUrl:%v",picUrl)
	return
	//dataArraryBytes,_:=dataArrary.Bytes()
	//dataArraryJson,_:= simplejson.NewJson(dataArraryBytes)
	//log.Printf("dataArraryJson:%s",dataArraryJson)

}
func SendPicWithKey(key string){
	picUrl,err:=GetPicUrlWithKey(key)
	if err!=nil{
		log.Printf("GetPicUrl函数出错:%v",err)
		miraiHttp.SendText(err.Error())
		return
	}
	//picUrl:="https://i.pixiv.cat/img-original/img/2014/04/05/23/26/27/42702642_p0.jpg"
	miraiHttp.SendPic(picUrl)
}
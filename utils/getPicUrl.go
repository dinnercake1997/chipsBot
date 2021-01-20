package utils

import (
	"errors"
	"log"

	"github.com/bitly/go-simplejson"
)

func GetPicUrl()(picUrl string,err error){
	url:="https://api.lolicon.app/setu/?r18=1&size1200=true&apikey=908789905fb629c064ccf8"
	content,err:=DoGet(url)
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
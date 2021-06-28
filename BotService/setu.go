package BotService

import (
	"chipsBot/miraiHttp"
	"chipsBot/utils"
	"errors"
	"github.com/bitly/go-simplejson"
	"log"
	"net/url"
)

var selector struct{

}

//从luolikonAPI获取图片Url
func GetPicUrl()(picUrl string,err error){
	url:="https://api.lolicon.app/setu/?r18=2&size1200=true&apikey=908789905fb629c064ccf8"
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
func SendPicByQQ(qq string){
	//picUrl:="https://i.pixiv.cat/img-original/img/2014/04/05/23/26/27/42702642_p0.jpg"
	picUrl,err:=GetPicUrl()
	if err!=nil{
		log.Printf("GetPicUrl函数出错:%v",err)
		miraiHttp.SendTextByQQ(err.Error(),qq)
		return
	}
	miraiHttp.SendPicByQQ(picUrl,qq)
}

//从luolikonAPI获取图片Url
func GetPicUrlWithKey(key string)(picUrl string,err error){
	getPicUrl:="https://api.lolicon.app/setu/?r18=2&size1200=true&apikey=908789905fb629c064ccf8"
	keyWord:="&keyword="+url.QueryEscape(key)
	getPicUrl=getPicUrl+keyWord
	log.Printf("请求loliconAPI的url为:%v",getPicUrl)
	content,err:=utils.DoGet(getPicUrl)
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
	if len(dataArrary)==0{
		log.Printf("请求loliconAPIkey获取图片为空")
		return "",errors.New("请求loliconAPIkey获取图片为空")
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
		miraiHttp.SendText("我也不知道出了什么问题，总之就怪你的XP系统有毒吧~(0v0)~")
		//miraiHttp.SendText(err.Error())
		return
	}
	//picUrl:="https://i.pixiv.cat/img-original/img/2014/04/05/23/26/27/42702642_p0.jpg"
	miraiHttp.SendPic(picUrl)
}
func SendPicWithKeyByQQ(key string,qq string){
	picUrl,err:=GetPicUrlWithKey(key)
	if err!=nil{
		log.Printf("GetPicUrl函数出错:%v",err)
		miraiHttp.SendTextByQQ("我也不知道出了什么问题，总之就怪你的XP系统有毒吧~(0v0)~",qq)
		//miraiHttp.SendText(err.Error())
		return
	}
	//picUrl:="https://i.pixiv.cat/img-original/img/2014/04/05/23/26/27/42702642_p0.jpg"
	miraiHttp.SendPicByQQ(picUrl,qq)
}

func SendPicWithKeyAndQQGroup(key string,qqGroup string){
	picUrl,err:=GetPicUrlWithKey(key)
	if err!=nil{
		log.Printf("GetPicUrl函数出错:%v",err)
		miraiHttp.SendTextByQQ("我也不知道出了什么问题，总之就怪你的XP系统有毒吧~(0v0)~",qqGroup)
		//miraiHttp.SendText(err.Error())
		return
	}
	//picUrl:="https://i.pixiv.cat/img-original/img/2014/04/05/23/26/27/42702642_p0.jpg"
	miraiHttp.SendPicByQQ(picUrl,qqGroup)
}
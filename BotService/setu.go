package BotService

import (
	"chipsBot/miraiHttp"
	"chipsBot/utils"
	json2 "encoding/json"
	"errors"
	"fmt"
	"github.com/bitly/go-simplejson"
	"log"
	"net/url"
	"strconv"
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
	//picUrl,ok=dataMap["url"].(string)
	pid:=dataMap["pid"].(json2.Number)
	p:=dataMap["p"].(json2.Number)
	pidInt64,_:=pid.Int64()
	pInt64,_:=p.Int64()
	pidString:=strconv.Itoa(int(pidInt64))
	//pageInt,_:=strconv.Atoi(page)
	//picUrl2,ok:=dataMap["url"].(string)
	picUrl=GeneratePicUrlByIdAndPageAndPageCount(pidString,int(pInt64),int(pInt64))
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
	log.Printf("dataMap:%v",dataMap)
	for k, v := range dataMap{
		fmt.Printf("key:%v value:%v\n", k, v)
	}
	pid:=dataMap["pid"].(json2.Number)
	p:=dataMap["p"].(json2.Number)
	pidInt64,_:=pid.Int64()
	pInt64,_:=p.Int64()
	pidString:=strconv.Itoa(int(pidInt64))
	//pageInt,_:=strconv.Atoi(page)
	//picUrl2,ok:=dataMap["url"].(string)
	picUrl=GeneratePicUrlByIdAndPageAndPageCount(pidString,int(pInt64),int(pInt64))

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

//从luolikonAPI获取图片Url
func GetPicUrlWithKeyAnyWay(key string)(picUrl string,picUrl2 string,err error){
	getPicUrl:="https://api.lolicon.app/setu/?r18=2&size1200=true&apikey=908789905fb629c064ccf8"
	keyWord:="&keyword="+url.QueryEscape(key)
	getPicUrl=getPicUrl+keyWord
	log.Printf("请求loliconAPI的url为:%v",getPicUrl)
	content,err:=utils.DoGet(getPicUrl)
	if err!=nil{
		log.Printf("请求loliconAPIkey出错:%v",err)
		return "","",errors.New("请求loliconAPIkey出错")
	}
	log.Printf("content:%v",content)
	json,err:= simplejson.NewJson([]byte(content))
	if err!=nil{
		log.Printf("请求loliconAPIkey转json出错:%v",err)
		return "","",errors.New("请求loliconAPIkey转json出错")
	}
	if err!=nil{return}
	dataArrary, err :=json.Get("data").Array()
	if err!=nil{
		log.Printf("请求loliconAPIkey获取图片信息出错:%v",err)
		return "","",errors.New("请求loliconAPIkey获取图片列表出错")
	}
	if len(dataArrary)==0{
		log.Printf("请求loliconAPIkey获取图片为空")
		return "","",errors.New("请求loliconAPIkey获取图片为空")
	}
	dataMap,ok:= dataArrary[0].(map[string]interface{})
	if !ok{
		log.Printf("请求loliconAPIkey获取图片url出错:%v",err)
		return "","",errors.New("请求loliconAPIkey获取图片信息出错")
	}
	log.Printf("dataMap:%v",dataMap)
	for k, v := range dataMap{
		fmt.Printf("key:%v value:%v\n", k, v)
	}
	pid:=dataMap["pid"].(json2.Number)
	p:=dataMap["p"].(json2.Number)
	pidInt64,_:=pid.Int64()
	pInt64,_:=p.Int64()
	pidString:=strconv.Itoa(int(pidInt64))
	//pageInt,_:=strconv.Atoi(page)
	//picUrl2,ok:=dataMap["url"].(string)
	picUrl=GeneratePicUrlByIdAndPageAndPageCountWithNoP(pidString,int(pInt64),int(pInt64))
	picUrl2=GeneratePicUrlByIdAndPageAndPageCountWithP(pidString,int(pInt64),int(pInt64))
	if !ok{
		log.Printf("请求loliconAPIkey获取图片url出错:%v",err)
		return "","",errors.New("请求loliconAPIkey获取图片url出错")
	}
	log.Printf("picUrl:%v",picUrl)
	return
	//dataArraryBytes,_:=dataArrary.Bytes()
	//dataArraryJson,_:= simplejson.NewJson(dataArraryBytes)
	//log.Printf("dataArraryJson:%s",dataArraryJson)

}
//从luolikonAPI获取图片Url
func GetPicUrlWithKeyWitoutModify(key string)(picUrl string,err error){
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
	log.Printf("dataMap:%v",dataMap)
	for k, v := range dataMap{
		fmt.Printf("key:%v value:%v\n", k, v)
	}
	//pid:=dataMap["pid"].(json2.Number)
	//p:=dataMap["p"].(json2.Number)
	//pidInt64,_:=pid.Int64()
	//pInt64,_:=p.Int64()
	//pidString:=strconv.Itoa(int(pidInt64))
	////pageInt,_:=strconv.Atoi(page)
	picUrl, ok = dataMap["url"].(string)
	//picUrl=GeneratePicUrlByIdAndPageAndPageCount(pidString,int(pInt64),int(pInt64))

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
		miraiHttp.SendText("欧尼酱的XP系统太过前卫了，阿比爱莫能助(X灬X)")
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
		miraiHttp.SendTextByQQ("欧尼酱的XP系统太过前卫了，阿比爱莫能助(X灬X)",qq)
		//miraiHttp.SendText(err.Error())
		return
	}
	//picUrl:="https://i.pixiv.cat/img-original/img/2014/04/05/23/26/27/42702642_p0.jpg"
	miraiHttp.SendPicByQQ(picUrl,qq)
}

func SendPicWithKeyAndQQGroup(key string,qqGroup string){
	picUrl,picUrl2,err:=GetPicUrlWithKeyAnyWay(key)
	if err!=nil{
		log.Printf("GetPicUrl函数出错:%v",err)
		miraiHttp.SendTextByQQ("欧尼酱的XP系统太过前卫了，阿比爱莫能助(X灬X)",qqGroup)
		//miraiHttp.SendText(err.Error())
		return
	}
	//picUrl:="https://i.pixiv.cat/img-original/img/2014/04/05/23/26/27/42702642_p0.jpg"
	miraiHttp.SendPicByQQ(picUrl,qqGroup)
	miraiHttp.SendPicByQQ(picUrl2,qqGroup)
}
func SendPicWithKeyAndQQGroup2(key string,qqGroup string){
	picUrl,err:=GetPicUrlWithKeyWitoutModify(key)
	if err!=nil{
		log.Printf("GetPicUrl函数出错:%v",err)
		miraiHttp.SendTextByQQ("欧尼酱的XP系统太过前卫了，阿比爱莫能助(X灬X)",qqGroup)
		//miraiHttp.SendText(err.Error())
		return
	}
	//picUrl:="https://i.pixiv.cat/img-original/img/2014/04/05/23/26/27/42702642_p0.jpg"
	miraiHttp.SendPicByQQ(picUrl,qqGroup)
}
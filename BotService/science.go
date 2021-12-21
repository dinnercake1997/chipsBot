package BotService

import (
	"chipsBot/miraiHttp"
	"errors"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)



func getScience() (content string,err error){
	content+="给小可爱们送上科学探索咨询，让我们一起探索宇宙和科学的真相吧！！！\\n"
	apiUrl := "http://api.tianapi.com/sicprobe/index?key=5c647563d4afcb93a78ab3dbc1e731c9&num=10"
	req, _ := http.NewRequest("GET", apiUrl, nil)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err!=nil{
		log.Printf("请求接口出错")
		err=errors.New("请求接口出错")
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	json,err:= simplejson.NewJson(body)
	if err!=nil{
		log.Printf("接口转json出错:%v",err)
		return "",errors.New("请求转json出错")
	}
	if err!=nil{return}
	dataArrary, err :=json.Get("newslist").Array()
	if err!=nil{
		log.Printf("请求信息出错:%v",err)
		return "",errors.New("请求列表出错")
	}
	for i,item :=range dataArrary {
		if i>10{
			break
		}
		dataMap,ok:= item.(map[string]interface{})
		if !ok{
			log.Printf("请求出错:%v",err)
			return "",errors.New("请求出错")
		}
		title,_:=dataMap["title"].(string)
		url,_:=dataMap["url"].(string)
		iString:=strconv.Itoa(i+1)
		content+=iString+".题目:"+title+" 薯条咨询直通车:"+url+"\\n"
	}
	return
}

func ScienceSend(){
	text,err:=getScience()

	if err!=nil{
		log.Printf("请求求微博热点接口获取图片信息出错:%v",err)
	}
	//url,err:=utils.GetPicUrl()
	//if err!=nil{
	//	log.Printf("请求luolikon接口获取图片信息出错:%v",err)
	//}
	url:=""
	miraiHttp.SendMix(text,url)
}
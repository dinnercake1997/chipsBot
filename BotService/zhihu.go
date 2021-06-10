
package BotService

import (
	"chipsBot/miraiHttp"
	"errors"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

)

func getZhiHuReSou() (content string,err error){
	content+="给大可爱们送上知乎实时热榜\\n"
	url := "http://api.rosysun.cn/zhihu/"
	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err!=nil{
		log.Printf("请求知乎热点接口出错")
		err=errors.New("请求知乎热点接口出错")
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	json,err:= simplejson.NewJson(body)
	if err!=nil{
		log.Printf("请求知乎热点接口转json出错:%v",err)
		return "",errors.New("请求知乎热点接口转json出错")
	}
	if err!=nil{return}
	dataArrary, err :=json.Get("data").Array()
	if err!=nil{
		log.Printf("请求知乎热点接口获取数组信息出错:%v",err)
		return "",errors.New("请求知乎热点接口获取数组列表出错")
	}
	for i,item :=range dataArrary {
		if i>10{
			break
		}
		dataMap,ok:= item.(map[string]interface{})
		if !ok{
			log.Printf("请求知乎出错:%v",err)
			return "",errors.New("请求出错")
		}
		title,_:=dataMap["title"].(string)
		_,_=dataMap["url"].(string)
		iString:=strconv.Itoa(i+1)
		content+=iString+".话题:"+title+" 传送门:"+"(暂时不给了)"+"\\n"
		//log.Printf("title:%v",title)
	}
	randInt := rand.Intn(100)  //生成0-99之间的随机数
	content+="带个随机数防止腾讯以为是广告："+string(randInt)
	return
}

func ZhiHuSend(){
	text,err:=getZhiHuReSou()

	if err!=nil{
		log.Printf("请求求微博热点接口获取图片信息出错:%v",err)
	}
	//url,err:=utils.GetPicUrl()
	//if err!=nil{
	//	log.Printf("请求luolikon接口获取图片信息出错:%v",err)
	//}
	url:="https://gitee.com/liang_zi_hao1997/FacemaskOrder/blob/master/weiboLogo.png"
	miraiHttp.SendMix(text,url)
}
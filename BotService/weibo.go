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

func getWeiboReSou() (content string,err error){
	content+="给小基佬们送上微博实时热搜\\n"
	url := "http://api.tianapi.com/txapi/weibohot/index?key=5c647563d4afcb93a78ab3dbc1e731c9"
	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err!=nil{
		log.Printf("请求微博热点接口出错")
		err=errors.New("请求微博热点接口出错")
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	json,err:= simplejson.NewJson(body)
	if err!=nil{
		log.Printf("请求求微博热点接口转json出错:%v",err)
		return "",errors.New("请求求微博热点接口转json出错")
	}
	if err!=nil{return}
	dataArrary, err :=json.Get("newslist").Array()
	if err!=nil{
		log.Printf("请求求微博热点接口获取图片信息出错:%v",err)
		return "",errors.New("请求求微博热点接口获取图片列表出错")
	}
	for i,item :=range dataArrary {
		if i>10{
			break
		}
		dataMap,ok:= item.(map[string]interface{})
		if !ok{
			log.Printf("请求loliconAPIkey获取图片url出错:%v",err)
			return "",errors.New("请求求微博热点接口获取图片信息出错")
		}
		hotword,_:=dataMap["hotword"].(string)
		hotwordnum,_:=dataMap["hotwordnum"].(string)
		iString:=strconv.Itoa(i+1)
		content+=iString+".热搜:"+hotword+" 热度:"+hotwordnum+"\\n"
	}
	randInt := rand.Intn(100)  //生成0-99之间的随机数
	content+="带个随机数防止腾讯以为是广告："+string(randInt)
	return
}

func WeiBoSend(){
	text,err:=getWeiboReSou()

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
func WeiBoSendByQQ(qq string){
	text,err:=getWeiboReSou()

	if err!=nil{
		log.Printf("请求求微博热点接口获取图片信息出错:%v",err)
	}
	//url,err:=utils.GetPicUrl()
	//if err!=nil{
	//	log.Printf("请求luolikon接口获取图片信息出错:%v",err)
	//}
	//url:="https://gitee.com/liang_zi_hao1997/FacemaskOrder/blob/master/weiboLogo.png"
	miraiHttp.SendMixByQQ(text,qq)
}
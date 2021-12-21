package PublicService

import (
	"errors"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func GetWeiboReSouPublic() (content string,err error){
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

	return
}
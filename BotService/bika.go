package BotService

import (
	"bytes"
	"chipsBot/config"
	"chipsBot/miraiHttp"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//func initBika()  {
//	header := map[string]string{
//		"api-key":           "C69BAF41DA5ABD1FFEDC6D2FEA56B",
//		"accept":            "application/vnd.picacomic.com.v1+json",
//		"app-channel":       "2",
//		"time":              strconv.FormatInt(time.Now().Unix(), 10),
//		"nonce":             "b1ab87b4800d4d4590a11701b8551afa",
//		"signature":         "",
//		"app-version":       "2.2.1.2.3.3",
//		"app-uuid":          "defaultUuid",
//		"app-platform":      "android",
//		"app-build-version": "44",
//		"Content-Type":      "application/json; charset=UTF-8",
//		"User-Agent":        "okhttp/3.8.1",
//		"authorization":     token,
//		"image-quality":     "original",
//	}
//
//
//}
var header map[string]string

var  tempToken  string
func InitBika(method string,url string){

	header = map[string]string{
		"api-key":           "C69BAF41DA5ABD1FFEDC6D2FEA56B",
		"accept":            "application/vnd.picacomic.com.v1+json",
		"app-channel":       "2",
		"time":              strconv.FormatInt(time.Now().Unix(), 10),
		"nonce":             "b1ab87b4800d4d4590a11701b8551afa",
		"signature":         "",
		"app-version":       "2.2.1.2.3.3",
		"app-uuid":          "defaultUuid",
		"app-platform":      "android",
		"app-build-version": "44",
		"Content-Type":      "application/json; charset=UTF-8",
		"User-Agent":        "okhttp/3.8.1",
		"authorization":    	tempToken,
		"image-quality":     "original",
	}
	if url=="" {return}
	//signInUrl:="https://picaapi.picacomic.com/auth/sign-in"
	var nonce = "b1ab87b4800d4d4590a11701b8551afa"
	var apiKey = "C69BAF41DA5ABD1FFEDC6D2FEA56B"
	//var raw = strings.Replace(signInUrl, "https://picaapi.picacomic.com/",    "", 1) + header["time"] + nonce + "POST" + apiKey
	//raw = strings.ToLower(raw)
	//h := hmac.New(sha256.New, []byte("~d}$Q7$eIni=V)9\\RK/P.RM4;9[7|@/CA}b~OW!3?EV`:<>M7pddUBL5n|0/*Cn"))
	//h.Write([]byte(raw))
	//header["signature"] = hex.EncodeToString(h.Sum(nil))
	//log.Println("header[\"signature\"]:"+header["signature"])
	//setToken()
	raw := strings.Replace(url, "https://picaapi.picacomic.com/",    "", 1) + header["time"] + nonce + method + apiKey
	raw = strings.ToLower(raw)
	h := hmac.New(sha256.New, []byte("~d}$Q7$eIni=V)9\\RK/P.RM4;9[7|@/CA}b~OW!3?EV`:<>M7pddUBL5n|0/*Cn"))
	h.Write([]byte(raw))
	log.Printf("time=%s",header["time"])

	header["signature"] = hex.EncodeToString(h.Sum(nil))
	log.Printf("signature=%s",header["signature"])
	//SetToken()
}
func SetBikaToken() {

	url:="https://picaapi.picacomic.com/auth/sign-in"
	InitBika("POST",url)
	requestBody:="{\"email\":\"dinnercake1997\",\"password\":\"qq1020224260\"}"
		log.Println("开始登录bika")
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		log.Printf("requestBody：%s",requestBody)
		jsonStr := []byte(requestBody)
		err:= writer.Close()
		if err != nil {
			return
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		if err != nil {
			log.Println("请求出错")

		}
		//token:="access_token="+accessToken
		//req.Header.Add("Cookie", token)
		for item := range header{
			req.Header.Add(item, header[item])
		}

		//req.Header.Add("", "")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("请求出错:"+err.Error())

		}
		log.Println(header)
		log.Println(resp.Body)
		bytes, err := ioutil.ReadAll(resp.Body)
		content:=string(bytes)
		log.Println(content)
		if err != nil {
			log.Println("读取响应体出错")

		}

		json,err:= simplejson.NewJson(bytes)
		log.Printf("json:%v",json)
		data:=json.Get("data")
		token,_:=data.Get("token").String()
		tempToken=token
		log.Printf("authorization:"+token)
		resp.Body.Close()


}

func test( ){

	//url:="https://picaapi.picacomic.com/comics"
	//param:=""
}

func GetCosPlay(keyword string ,sort string,page string)(picUrl string,err error){
	bookId,err:=SearchComic("Cosplay",keyword,sort,page)
	if err!=nil{return}
	picUrl,err=enterBookAndSelectPaper(bookId)
	if err!=nil{return}
	return
}
func SearchLikePic(page string)(picUrl string,err error){
	rand.Seed(time.Now().UnixNano())
	url:= "https://picaapi.picacomic.com/users/favourite?s=dd&page="+page
	param0:=map[string]string{}
	res,err:=BikaActionGet(url,param0)
	if err!=nil{
		return
	}
	data:=res.Get("data")
	comics:=data.Get("comics")
	total,_:=comics.Get("total").Int()
	if total==0{
		return "",errors.New("没有找到cos或者服务器出错了")
	}
	var take int
	if total>20{take=rand.Intn(19)}else{take=rand.Intn(total-1)}
	log.Printf("抽到第几本？%s",take)
	docsArray,_:=comics.Get("docs").Array()
	var bookId string
	for i,item :=range docsArray {
		if i==take{
			dataMap,ok:= item.(map[string]interface{})
			if !ok{
				log.Printf("请求cos的docs出错:%v",err)
				return "",errors.New("请求cos的docs出错")
			}//找到哪本了cos
			bookId,_=dataMap["_id"].(string)
			name,_:=dataMap["title"].(string)
			log.Printf("抽到的bookId为:%v",bookId)
			log.Printf("抽到的bookName为:%v",name)
		}
	}

	//进入漫画
	rand.Seed(time.Now().UnixNano())
	bookUrl:="https://picaapi.picacomic.com/comics/"+bookId+"/order/1/pages?page=1"
	//找由大页
	var param map[string]string
	res,err=BikaActionGet(bookUrl,param)
	if err!=nil {return "",err}
	bookData:=res.Get("data")
	pageData:=bookData.Get("pages")
	docs,err:=pageData.Get("docs").Array()
	var takePic int
	log.Printf("num:%v",len(docs))
	if len(docs)==0{
		return "",errors.New("num为0")
	}else{
		takePic= 1+rand.Intn(len(docs)-1)
	}

	if err!=nil {return "",err}
	for i,item :=range docs {
		if i==takePic-1{
			dataMap,ok:= item.(map[string]interface{})
			if !ok{
				log.Printf("请求cos的docs出错:%v",err)
				return "",errors.New("请求cos的docs出错")
			}//找到哪本了cos
			paperId,_:=dataMap["_id"].(string)
			media,_:=dataMap["media"]
			mediaMap,ok:=media.(map[string]interface{})
			path:=mediaMap["path"].(string)
			fileServer:=mediaMap["fileServer"].(string)
			log.Printf("抽到的paperId为:%v",paperId)
			picUrl=fileServer+"/static/"+path
			log.Printf("抽到的picUrl为:%v",picUrl)
			return picUrl,nil
		}
	}
	log.Printf("啥都没有")
	return picUrl,errors.New("啥都没有")
}

func SendCosWithQQGroup(qqGroup string) (err error){
	rand.Seed(time.Now().UnixNano())
	tekePages:=1+rand.Intn(int(config.Myconfig.BikaStarPagesCount))
	tekePagesString:=strconv.Itoa(tekePages)
	picUrl,err:=SearchLikePic(tekePagesString)
	if err!=nil{
		miraiHttp.SendTextByQQ(err.Error(),qqGroup)
		return
	}
	miraiHttp.SendPicByQQ(picUrl,qqGroup)
	return
}
func SearchComic(categories string,keyword string ,sort string,page string)(bookId string,err error){
	rand.Seed(time.Now().UnixNano())
	param:="{\"categories\":[\""+categories+"\"],\"keyword\":\""+keyword+"\",\"sort\":\""+sort+"\"}"
	url:= "https://picaapi.picacomic.com/comics/advanced-search?page="+page
	res,err:=BikaActionPost(url,param)
	if err!=nil{
		return
	}
	data:=res.Get("data")
	comics:=data.Get("comics")
	total,_:=comics.Get("total").Int()
	if total==0{
		return "",errors.New("没有找到cos或者服务器出错了")
	}
	var take int
	if total>20{take=rand.Intn(19)}else{take=rand.Intn(total-1)}
	log.Printf("抽到第几本？%s",take)
	docsArray,_:=comics.Get("docs").Array()
	for i,item :=range docsArray {
		if i==take{
			dataMap,ok:= item.(map[string]interface{})
			if !ok{
				log.Printf("请求cos的docs出错:%v",err)
				return "",errors.New("请求cos的docs出错")
			}//找到哪本了cos
			bookId,_=dataMap["_id"].(string)
			name,_:=dataMap["title"].(string)
			log.Printf("抽到的bookId为:%v",bookId)
			log.Printf("抽到的bookName为:%v",name)
		}
	}
	return
}
func enterBookAndSelectPaper(bookId string)(picUrl string,err error){
	rand.Seed(time.Now().UnixNano())
	bookUrl:="https://picaapi.picacomic.com/comics/"+bookId+"/order/1/pages?page=1"
	//找由大页
	var param map[string]string
	res,err:=BikaActionGet(bookUrl,param)
	if err!=nil {return "",err}
	data:=res.Get("data")
	docs,err:=data.Get("docs").Array()
	num :=0
	var takePic int
	for _,_ =range docs {
		num++
	}
	log.Printf("num:%v",num)
	if num==0{
		return "",errors.New("num为0")
	}else{
		takePic= 1+rand.Intn(num-1)
	}

	if err!=nil {return "",err}
	for i,item :=range docs {
		if i==takePic-1{
			dataMap,ok:= item.(map[string]interface{})
			if !ok{
				log.Printf("请求cos的docs出错:%v",err)
				return "",errors.New("请求cos的docs出错")
			}//找到哪本了cos
			paperId,_:=dataMap["_id"].(string)
			log.Printf("抽到的paperId为:%v",paperId)
			picUrl="https://storage1.picacomic.com/static/"+paperId+".jpg"
			log.Printf("抽到的picUrl为:%v",picUrl)
			return picUrl,nil
		}
	}
	log.Printf("啥都没有")
	return picUrl,errors.New("啥都没有")
}





func BikaActionPost(url string ,param string ) (res *simplejson.Json,err error) {

	log.Println("开始请求BikaActionPost")
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	//requestBody:="{\"email\":\"qq1020224260\",\"password\":\"qq1020224260\"}"
	//log.Printf("requestBody：%s",requestBody)
	log.Printf("requestBody：%s",param)
	jsonStr := []byte(param)
	err= writer.Close()
	if err != nil {
		return
	}
	InitBika("POST",url)
	var req *http.Request
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("请求出错")
	}
	for item := range header{
		req.Header.Add(item, header[item])
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("请求出错:"+err.Error())
		return  res,errors.New("请求出错")
	}
	log.Println(resp.Body)
	bytes, err := ioutil.ReadAll(resp.Body)
	//content := string(bytes)
	if err != nil {
		log.Println("读取响应体出错")

	}

	res, err = simplejson.NewJson(bytes)
	//log.Printf("res:%s",content)
	return
}


func BikaActionGet(getUrl string ,param map[string]string) (res *simplejson.Json,err error) {


	log.Println("开始请求")
	log.Printf("requestUrl：%s",getUrl)
	var req *http.Request
	getUrl+="?"
	for item :=range param{
		escapeParam := url.QueryEscape(param[item])
		getUrl+=item+"="+escapeParam+"&"
	}
	if find := strings.Contains(getUrl, "&"); find {
		getUrl=getUrl[:len(getUrl)-1]
	}
	//escapeUrl := url.QueryEscape(getUrl)
	req, err = http.NewRequest("GET", getUrl, nil)
	if err != nil {
		log.Println("请求出错")
		return
	}

	log.Printf("url:%s",getUrl)
	InitBika("GET",getUrl)
	log.Println("InitBika完成")
	for item := range header{
		req.Header.Add(item, header[item])
		//log.Printf("key:%S;value:%s",item,header[item])
	}
	log.Printf("header:%v;",req.Header)
	client:=&http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Printf("do出错:%s",err)
		return
	}
	bodyRead, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("read出错")
		return
	}

	if err != nil {
		log.Println("读取响应体出错")
		return

	}
	res, err = simplejson.NewJson(bodyRead)
	if err != nil {
		log.Println("simplejson.NewJson(bodyRead)出错")
		return
	}
	//content := string(bodyRead)
	//log.Printf("content:%s",content)
	return
}

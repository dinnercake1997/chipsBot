package main

import (
	"chipsBot/chat2bot"
	"chipsBot/config"
	"chipsBot/cronTask"
	"chipsBot/miraiHttp"
	"fmt"
	"log"
	"net/http"
)

func main(){

	//for{
	//	utils.SendPic()
	//	log.Printf("每5分钟打印一次\n")
	//	time.Sleep(time.Second*60*15)
	//}
	config.Myconfig=*config.Myconfig.GetConf()
	miraiHttp.InitMiraiHttp()
	cronTask.CronTask=cronTask.InitCronTask()
	//BotService.SendPic()
	//BotService.ScienceSend()
	//BotService.WeiBoSend()
	http.HandleFunc("/", Chat2Bot)
	http.HandleFunc("/message", TestHandler)
	err:=http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		// 服务器创建失败
		panic("服务器创建失败")
	}
}

func Chat2Bot(w http.ResponseWriter, r *http.Request) {
	// 获取请求报文的内容长度
	len := r.ContentLength
	// 新建一个字节切片，长度与请求报文的内容长度相同
	body := make([]byte, len)
	// 读取 r 的请求主体，并将具体内容读入 body 中
	r.Body.Read(body)
	log.Printf("body:%s",body)
	text,err:=chat2bot.GetTextFromBody(string(body))
	if err!=nil{
		log.Printf("GetTextFromBodyerr:%s",err)
	}
	chat2bot.ActionSelect(text)
	// 将字节切片内容写入相应报文
	fmt.Fprintln(w, body)
	return
}
func TestHandler(w http.ResponseWriter, r *http.Request) {
	// 获取请求报文的内容长度
	len := r.ContentLength
	// 新建一个字节切片，长度与请求报文的内容长度相同
	body := make([]byte, len)

	// 读取 r 的请求主体，并将具体内容读入 body 中
	r.Body.Read(body)
	text,err:=chat2bot.GetTextFromBody(string(body))
	if err!=nil{
		log.Printf("GetTextFromBodyerr:%s",err)
	}
	chat2bot.ActionSelect(text)
	log.Printf("body:%s",body)
	// 将字节切片内容写入相应报文
	fmt.Fprintln(w, body)

}

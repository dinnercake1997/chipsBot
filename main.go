package main

import (
	"chipsBot/BotService"
	"chipsBot/chat2bot"
	"chipsBot/config"
	"chipsBot/cronTask"
	"chipsBot/miraiHttp"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/NateScarlet/pixiv/pkg/artwork"
	"github.com/NateScarlet/pixiv/pkg/client"
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
	BotService.InitPixivAPI()

	//BotService.SendDailyPic()
	//test()
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


func test(){
	//client.Default
	//使用 PHPSESSID Cookie 登录 (推荐)。
	c := &client.Client{}
	c.BypassSNIBlocking()
	c.SetDefaultHeader("User-Agent", client.DefaultUserAgent)

	c.SetPHPSESSID("11517896_kmpPv4CJsn0ly8L1q0z8ZJX4LO100YpG")


	//c := &client.Client{}
	//c.SetDefaultHeader("User-Agent", client.DefaultUserAgent)
	//c.Login("1020224260@qq.com", "qq1020224260")

	var ctx = context.Background()
	ctx = client.With(ctx, c)


	// 搜索画作
	result, err := artwork.Search(ctx, "パチュリー・ノーレッジ")
	if err==nil{
		log.Printf("resultJson:%v",result.JSON)
		//result.Artworks() // []artwork.Artwork，只有部分数据，通过 `Fetch` `FetchPages` 方法获取完整数据。
		//artwork.Search(ctx, "パチュリー・ノーレッジ", artwork.SearchOptionPage(2)) // 获取第二页
	}else{
		log.Printf("err:%v",err)
	}



	// 画作排行榜
	rank := &artwork.Rank{Mode: "weekly_r18"}
	rank.Fetch(ctx)
	log.Printf("rankLen:%v",len(rank.Items))
	log.Printf("rank:%v",rank.Items[0].Rank)
	if  len(rank.Items)!=0{
		r := rank.Items[0].Rank
		//_=rank.Items[0].PreviousRank
		//_=rank.Items[0].Artwork
		log.Printf("rank:%v",r)
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

package BotService

import (
	"chipsBot/miraiHttp"
	"context"
	"errors"
	"github.com/NateScarlet/pixiv/pkg/artwork"
	"github.com/NateScarlet/pixiv/pkg/client"
	"log"
	"strings"
)

var c *client.Client
func getDailyPic() ([]artwork.RankItem,error){
	var ctx = context.Background()
	ctx = client.With(ctx, c)
	// 画作排行榜
	rank := &artwork.Rank{Mode: "weekly_r18"}
	rank.Fetch(ctx)
	log.Printf("rankLen:%v",len(rank.Items))
	log.Printf("rank:%v",rank.Items[0].Rank)
	if  len(rank.Items)!=0{
		//r := rank.Items[0].Rank
		//_=rank.Items[0].PreviousRank
		//_=rank.Items[0].Artwork
		//log.Printf("rank:%v",r)
		return rank.Items,nil
	}
	return nil,errors.New("搜索日榜为空！")
}
func SendDailyPic() error{
	items,err:=getDailyPic()
	if err!=nil{
		miraiHttp.SendText("")
	}
	for  i:=0;i<len(items)&&i<5;i++{
		url:=items[i].Image.Regular
		url=strings.Replace(url,"pximg.net/c/240x480","pixiv.cat",1)
		log.Printf("url:%s",url)
		miraiHttp.SendPic(url)
	}
	return nil
}
func InitPixivAPI(){
	//使用 PHPSESSID Cookie 登录 (推荐)。
	c = &client.Client{}
	c.BypassSNIBlocking()
	c.SetDefaultHeader("User-Agent", client.DefaultUserAgent)

	c.SetPHPSESSID("11517896_kmpPv4CJsn0ly8L1q0z8ZJX4LO100YpG")


	//c := &client.Client{}
	//c.SetDefaultHeader("User-Agent", client.DefaultUserAgent)
	//c.Login("1020224260@qq.com", "qq1020224260")


}
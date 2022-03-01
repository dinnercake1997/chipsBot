package BotService

import (
	"chipsBot/miraiHttp"
	"context"
	"errors"
	"github.com/NateScarlet/pixiv/pkg/artwork"
	"github.com/NateScarlet/pixiv/pkg/client"
	"log"
	"strconv"
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
func getDailyPic2() ([]artwork.RankItem,error){
	var ctx = context.Background()
	ctx = client.With(ctx, c)
	// 画作排行榜
	rank := &artwork.Rank{Mode: "daily"}
	rank.Fetch(ctx)
	log.Printf("rank:%v",rank)
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

func GetDailyPicByMode(mode string) ([]artwork.RankItem,error){

	var ctx = context.Background()
	ctx = client.With(ctx, c)
	// 画作排行榜
	rank := &artwork.Rank{Mode: mode}
	rank.Fetch(ctx)
	log.Printf("rankLen:%v",len(rank.Items))

	if  len(rank.Items)!=0{
		//r := rank.Items[0].Rank
		//_=rank.Items[0].PreviousRank
		//_=rank.Items[0].Artwork
		//log.Printf("rank:%v",r)
		log.Printf("rank:%v",rank.Items[0].Rank)
		return rank.Items,nil

	}

	return nil,errors.New("搜索日榜为空！")
}

func GetPixivPicByKey(key string) (url string,err error){
	var ctx = context.Background()
	ctx = client.With(ctx, c)
	log.Printf("开始搜索:%v",key)
	result, err := artwork.Search(ctx, key)
	if err!=nil{
		return "", err
	}
	//result.JSON
	res:=result.Artworks()
	log.Printf("resLen:%v", len(res))
	log.Printf("res1:%v,PageCount:%v,likeCount:%v",res[0].ID,res[0].PageCount)

	url=GeneratePicUrlByIdAndPageAndPageCount(res[0].ID, 1,int(res[0].PageCount))
	return

	//artwork.Search(ctx, key, artwork.SearchOptionPage(2)) // 获取第二页
	//res=result.Artworks()
	//log.Printf("res2:%v,pages:%v",res[0].ID,res[0].Pages)
	return "",errors.New("搜索为空！")
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

func SendDailyPicWithQQGroup(qqGroup string) error{
	items,err:=getDailyPic()
	if err!=nil{
		miraiHttp.SendText("")
	}
	for  i:=0;i<len(items)&&i<5;i++{
		url:=items[i].Image.Regular
		url=strings.Replace(url,"pximg.net/c/240x480","pixiv.cat",1)
		log.Printf("url:%s",url)
		miraiHttp.SendPicByQQ(url,qqGroup)
	}
	return nil
}
func SendDailyPicWithQQGroupByMode(qqGroup string,mode string) error{

	items,err:= GetDailyPicByMode(mode)
	if err!=nil{
		miraiHttp.SendTextByQQ(err.Error(),qqGroup)
	}
	for  i:=0;i<len(items)&&i<7;i++{
		pid:=items[i].ID
		count:=items[i].PageCount
		for j:=1;j<=int(count)&&j<=3;j++{
			url:=GeneratePicUrlByIdAndPageAndPageCount(pid,j,int(count))
			log.Printf("url:%s",url)
			miraiHttp.SendPicByQQ(url,qqGroup)
		}
		//url:=items[i].Image.Regular
		//url=strings.Replace(url,"pximg.net/c/240x480","pixiv.cat",1)
		//log.Printf("url:%s",url)
		//miraiHttp.SendPicByQQ(url,qqGroup)
	}
	return nil
}

func SendDailyPic2() error{
	items,err:=getDailyPic2()
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

func SendDailyPic2WithQQGroup(qqGroup string) error{
	items,err:=getDailyPic2()
	if err!=nil{
		miraiHttp.SendText("")
	}
	for  i:=0;i<len(items)&&i<5;i++{
		url:=items[i].Image.Regular
		//url=strings.Replace(url,"pximg.net/c/240x480","pixiv.cat",1)
		log.Printf("url:%s",url)
		miraiHttp.SendPicByQQ(url,qqGroup)
	}
	return nil
}
func InitPixivAPI(){

	//使用 PHPSESSID Cookie 登录 (推荐)。
	log.Printf("开始初始化pixiv接口！")
	c = &client.Client{}
	//c.BypassSNIBlocking()
	//c.SetDefaultHeader("User-Agent", client.DefaultUserAgent)
	//11517896_kmpPv4CJsn0ly8L1q0z8ZJX4LO100YpG
	c.SetPHPSESSID("11517896_d8PaKCTeRqKDc2SpRx835nVtFdHEnjd3")
	c.SetDefaultHeader("User-Agent", client.DefaultUserAgent)
	//c.Login("1020224260@qq.com", "qq1020224260")
	log.Printf("初始化pixiv接口完成！")

	//c := &client.Client{}
	//c.SetDefaultHeader("User-Agent", client.DefaultUserAgent)
	//c.Login("1020224260@qq.com", "qq1020224260")


}

func TestGetPixiv(){
	GetDailyPicByMode("daily")
}
func GeneratePicUrlByIdAndPageAndPageCount(id string,page int,pageCount int )(url string){
	p:= strconv.Itoa(page)
	if pageCount==1||pageCount==0{
		url="https://pixiv.re/"+id+".png"
	}else{
		url="https://pixiv.re/"+id+"-"+p+".png"
	}
	return url
}

func GeneratePicUrlByIdAndPageAndPageCountWithP(id string,page int,pageCount int )(url string){
	p:= strconv.Itoa(page)
		url="https://pixiv.re/"+id+"-"+p+".png"
	return url
}
func GeneratePicUrlByIdAndPageAndPageCountWithNoP(id string,page int,pageCount int )(url string){
	url="https://pixiv.re/"+id+".png"
	return url
}
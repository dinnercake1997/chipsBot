package chat2bot

import (
	"chipsBot/BotService"
	"chipsBot/miraiHttp"
	"errors"
	"log"
	"regexp"
)
func ActionSelect(text string)(err error)  {



	reg:=regexp.MustCompile("[来给求][点张份](?P<keyword>.*?)?[色涩瑟][图]")
	key:=reg.FindSubmatch([]byte(text))
	log.Printf("ActionSelect text:%v\n",text)
	if text=="今日榜单"{
		miraiHttp.SendText("今日份r18榜单是吧，等会就来，你先把裤子脱好吧")
		BotService.SendDailyPic()
		return
	}
	if text=="知乎热榜"{
		miraiHttp.SendText("你怎么不看瑟图看这些，我知道了，是贤者模式！")
		BotService.ZhiHuSend()
		return
	}
	if text=="微博热搜"{
		miraiHttp.SendText("你会觉得厕所里面会有什么好东西？姥爷你真是独具慧眼啊！")
		BotService.WeiBoSend()
		return
	}
	if text=="link start"{
		//miraiHttp.SendText("现在的功能还很少呢~大伙快想想还有什么好idea呀！")
		caidan:="1.来 点/份 XX 色/瑟图：XX为关键字，如果没有，则随机瑟图（别看那么多，小心群主被抓走！）\\n" +
			"2.今日榜单:p站r18日榜，只截取前5张.\\n" +
			"3.知乎热榜：懂的都懂。\\n" +
			"4.微博热搜：没啥好说的。"
		url:="https://gitee.com/liang_zi_hao1997/FacemaskOrder/blob/master/weiboLogo.png"
		miraiHttp.SendMix(caidan,url)
		return

	}
	if len(key)==0{
		return nil
	}
	if len(key[0])==0{
		//没有识别出
		return nil
	}
	if len(key[1])==0{
		BotService.SendPic()
	}else{

		log.Printf("ActionSelectKey1:%v\n",string(key[1]))
		BotService.SendPicWithKey(string(key[1]))
	}
	return errors.New("出错啦!")
}
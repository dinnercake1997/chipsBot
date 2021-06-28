package chat2bot

import (
	"chipsBot/BotService"
	"chipsBot/config"
	"chipsBot/miraiHttp"
	"chipsBot/utils"
	"errors"
	"log"
	"regexp"
	"strconv"
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
	if text=="今日女性榜单"{
		miraiHttp.SendText("今日份女性榜单是吧，等会就来，你先把裤子脱好吧")
		BotService.SendDailyPic2()
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


func ActionSelectWithQQGroup(text string,qqGroup string)(err error)  {
	isWhite:=false
	//qqSendGroups:=[]string{"1085171553","763091038"}
	for _, v := range config.QQSendGroups {
		if v==qqGroup{
			isWhite =true
			break
		}
	}
	if !isWhite{
		return
	}
	reg:=regexp.MustCompile("[来给求][点张份](?P<keyword>.*?)?[色涩瑟][图]")
	key:=reg.FindSubmatch([]byte(text))
	log.Printf("ActionSelect text:%v\n",text)


	if text=="今周榜单"{
		miraiHttp.SendTextByQQ("要去了是吧，你先把裤子脱好吧",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"weekly")
		return
	}
	if text=="今日榜单"{
		miraiHttp.SendTextByQQ("要去了是吧，你先把裤子脱好吧",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"daily")
		return
	}
	if text=="今月榜单"{
		miraiHttp.SendTextByQQ("要去了是吧，你先把裤子脱好吧",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"monthly")
		return
	}
	if text=="今周r18榜单"{
		miraiHttp.SendTextByQQ("要去了是吧，你先把裤子脱好吧",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"weekly_r18")
		return
	}
	if text=="今日r18榜单"{
		miraiHttp.SendTextByQQ("要去了是吧，你先把裤子脱好吧",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"daily_r18")
		return
	}
	if text=="今日男性r18榜单"{
		miraiHttp.SendTextByQQ("要去了是吧，你先把裤子脱好吧",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"male_r18")
		return
	}
	if text=="今日女性r18榜单"{
		miraiHttp.SendTextByQQ("要去了是吧，你先把裤子脱好吧",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"female_r18")
		return
	}
	if text=="今日奇怪r18榜单"{
		miraiHttp.SendTextByQQ("要去了是吧，你先把裤子脱好吧",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"r18g")
		return
	}



	if text=="知乎热榜"{
		miraiHttp.SendTextByQQ("你怎么不看瑟图看这些，我知道了，是贤者模式！",qqGroup)
		BotService.ZhiHuSendByQQ(qqGroup)
		return
	}
	if text=="微博热搜"{
		miraiHttp.SendTextByQQ("你会觉得厕所里面会有什么好东西？姥爷你真是独具慧眼啊！",qqGroup)
		BotService.WeiBoSendByQQ(qqGroup)
		return
	}
	if text=="开启定时色图"{
		config.SendOnTitleMap[qqGroup]="默认"
		miraiHttp.SendTextByQQ("开启整点定时通知施法~",qqGroup)
		return
	}
	if text=="关闭定时色图"{
		config.SendOnTitleMap[qqGroup]=""
		miraiHttp.SendTextByQQ("关闭定时色图成功（7 v 7 ）",qqGroup)
		return
	}
	if text=="重置主题"{
		config.SendOnTitleMap[qqGroup]="默认"
		miraiHttp.SendTextByQQ("重置主题成功（8 v 8）",qqGroup)
		return
	}


	if text=="link start"{
		//miraiHttp.SendText("现在的功能还很少呢~大伙快想想还有什么好idea呀！")
		caidan:="1.来 点/份 XX 色/瑟图：XX为关键字（“来点御姐瑟图”）\\n" +
			"2.榜单功能:p站r18日榜，只截取前5张.\\n" +
			"      “今日/周/月榜单”\\n" +
			"      ”今日/周r18榜单“\\n" +
			"      “今日男/女性r18榜单”\\n" +
			"3.知乎热榜：“知乎热榜”。\\n" +
			"4.微博热搜：“微博热搜”。\\n" +
			"5.定时瑟图：“开启定时色图” “关闭定时色图” “设置白丝主题” “重置主题” ”设置黑丝主题“等。\\n"+
			"6.配菜功能：“上配菜” 根据主题推送若干张图 。\\n"

		//url:="https://gitee.com/liang_zi_hao1997/FacemaskOrder/blob/master/weiboLogo.png"
		miraiHttp.SendMixByQQ(caidan,qqGroup)
		return

	}
	if text=="上配菜" {
		num:=utils.GetRand(1,5)
		numString:=strconv.Itoa(num)
		miraiHttp.SendTextByQQ("给你找到了"+numString+"张，看几张就给爷坚持几小时（= v = ）",qqGroup)
		if config.SendOnTitleMap[qqGroup]==""||config.SendOnTitleMap[qqGroup]=="默认"{

			for i:=0;i<num;i++{
				BotService.SendPicByQQ(qqGroup)
			}
		}else {
			for i:=0;i<num;i++{
				BotService.SendPicWithKeyAndQQGroup(config.SendOnTitleMap[qqGroup],qqGroup)
			}

		}
	}
	regTitle:=regexp.MustCompile("[设][置](?P<keyword>.*?)?[主][题]")
	keyTittle:=regTitle.FindSubmatch([]byte(text))

	if len(keyTittle)!=0&&len(keyTittle[1])!=0{
		config.SendOnTitleMap[qqGroup]=string(keyTittle[1])
		sendText:="设置主题"+string(keyTittle[1])+"成功辣!不过图库有没有你想要的我就不知道辣！~"
		miraiHttp.SendTextByQQ(sendText,qqGroup)
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
		if config.SendOnTitleMap[qqGroup]==""||config.SendOnTitleMap[qqGroup]=="默认"{
			BotService.SendPicByQQ(qqGroup)
		}else {
			BotService.SendPicWithKeyAndQQGroup(config.SendOnTitleMap[qqGroup],qqGroup)
		}

	}else{

		log.Printf("ActionSelectKey1:%v\n",string(key[1]))
		BotService.SendPicWithKeyAndQQGroup(string(key[1]),qqGroup)
	}
	return errors.New("出错啦!")
}

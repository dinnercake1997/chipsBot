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
		miraiHttp.SendTextByQQ("(*/_＼*)主人要那个是吗，需要我帮忙吗？我的意思是，需要我给你去门口把关吗？",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"weekly")
		return
	}
	if text=="今日榜单"{
		miraiHttp.SendTextByQQ("如果主人想看更多的，你只要和我说就行的，我也会一点，一点绘画技巧",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"daily")
		return
	}
	if text=="今月榜单"{
		miraiHttp.SendTextByQQ("(#／。＼#)主人，这次完事之后记得。。。记得穿好裤子",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"monthly")
		return
	}
	if text=="今周r18榜单"{
		miraiHttp.SendTextByQQ("ヽ(*。>Д<)o゜今晚我会晚点回来的，牛鞭，熊掌，肾宝，得好好慰劳疲惫的主人才行！",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"weekly_r18")
		return
	}
	if text=="今日r18榜单"{
		miraiHttp.SendTextByQQ("(*/_＼*)主人要那个是吗，需要我帮忙吗？我的意思是，需要我给你去门口把关吗",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"daily_r18")
		return
	}
	if text=="今日男性r18榜单"{
		miraiHttp.SendTextByQQ("(✿◡‿◡)主人正值少年，有这种需求，诺艾尔理解的。",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"male_r18")
		return
	}
	if text=="今日女性r18榜单"{
		miraiHttp.SendTextByQQ("(((φ(◎ロ◎;)φ)))不要给我看这些辣！对我来说还太早了",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"female_r18")
		return
	}
	if text=="今日奇怪r18榜单"{
		miraiHttp.SendTextByQQ("(o゜▽゜)o☆主人这么做，一定有其中的深意吧",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"r18g")
		return
	}



	if text=="知乎热榜"{
		miraiHttp.SendTextByQQ("(*･v･)知行合一，主人一直是我的榜样。",qqGroup)
		BotService.ZhiHuSendByQQ(qqGroup)
		return
	}
	if text=="微博热搜"{
		miraiHttp.SendTextByQQ("o(*^▽^*)┛了解时事的主人，诺艾尔，喜欢~",qqGroup)
		BotService.WeiBoSendByQQ(qqGroup)
		return
	}
	if text=="开启定时色图"{
		config.GroupSets[qqGroup].IsOnTime=true
		miraiHttp.SendTextByQQ("好的主人，我将会每个小时来为您送上配菜。\\n（( >ρ < ”)碎碎念:身体要是那么好，别浪费在那种地方啊",qqGroup)
		return
	}
	if text=="关闭定时色图"{
		config.GroupSets[qqGroup].IsOnTime=false
		miraiHttp.SendTextByQQ("┭┮﹏┭┮主人，你是身体不行了吗！？诺艾尔很担心你！",qqGroup)
		return
	}
	if text=="重置主题"{
		config.GroupSets[qqGroup].Tittle="随机"
		miraiHttp.SendTextByQQ("为主人重装xp系统成功！",qqGroup)
		return
	}
	if text=="查看主题"{
		tittle:=config.GroupSets[qqGroup].Tittle
		miraiHttp.SendTextByQQ("依我所知，主人现在的性癖是："+tittle,qqGroup)
		return
	}

	if text=="link start"{
		//miraiHttp.SendText("现在的功能还很少呢~大伙快想想还有什么好idea呀！")
		caidan:=

			"1.点图功能：“来点jk瑟图” “来点瑟图” “来点明日方舟瑟图” 可发挥想象力找想看的\\n\\n" +
			"2.榜单功能:\\n" +
			"      “今日/周/月榜单”\\n" +
			"      ”今日/周r18榜单“\\n" +
			"      “今日男/女性r18榜单”\\n\\n" +
			"3.新闻相关：“知乎热榜” “微博热搜”\\n\\n" +
			"4.定时瑟图：“开启定时色图” “关闭定时色图” “设置白丝主题” “重置主题” “查看主题” ”设置黑丝主题“等。\\n"+
			"5.配菜功能：“查看主题” “设置黑丝主题” “上配菜” 根据主题推送若干张图 。\\n"

		//url:="https://gitee.com/liang_zi_hao1997/FacemaskOrder/blob/master/weiboLogo.png"
		miraiHttp.SendMixByQQ(caidan,qqGroup)
		return

	}
	if text=="上配菜" {
		num:=utils.GetRand(1,5)
		numString:=strconv.Itoa(num)
		miraiHttp.SendTextByQQ("诺艾尔为主人准备了"+numString+"道菜",qqGroup)
		if config.GroupSets[qqGroup].Tittle=="随机"{
			for i:=0;i<num;i++{
				BotService.SendPicByQQ(qqGroup)
			}
		}else {
			for i:=0;i<num;i++{
				BotService.SendPicWithKeyAndQQGroup(config.GroupSets[qqGroup].Tittle,qqGroup)
			}

		}
	}
	regTitle:=regexp.MustCompile("[设][置](?P<keyword>.*?)?[主][题]")
	keyTittle:=regTitle.FindSubmatch([]byte(text))

	if len(keyTittle)!=0&&len(keyTittle[1])!=0{
		config.GroupSets[qqGroup].Tittle=string(keyTittle[1])
		sendText:="主人的性癖"+string(keyTittle[1])+"安装成功辣!诺艾尔会尽可能满足主人~"
		if string(keyTittle[1])=="诺艾尔"{
			sendText="(///_///)主人，请不要开这种玩笑！（传来脱衣解带的声音）"
			miraiHttp.SendTextByQQ(sendText,qqGroup)
			sendText="(#。﹏。#)请对诺艾尔温柔点。"

		}

		if string(keyTittle[1])=="女仆"{
			sendText="( >ρ < ”)你的眼前不就有一个女仆吗！"
		}

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
		BotService.SendPicByQQ(qqGroup)
	}else{
		log.Printf("ActionSelectKey1:%v\n",string(key[1]))
		BotService.SendPicWithKeyAndQQGroup(string(key[1]),qqGroup)
		return
	}
	return errors.New("出错啦!")
}

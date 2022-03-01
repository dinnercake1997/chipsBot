package chat2bot

import (
	"chipsBot/BotService"
	"chipsBot/config"
	"chipsBot/miraiHttp"
	"chipsBot/utils"
	"errors"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)
//func ActionSelect(text string)(err error)  {
//
//
//
//	reg:=regexp.MustCompile("[来给求][点张份](?P<keyword>.*?)?[色涩瑟][图]")
//	key:=reg.FindSubmatch([]byte(text))
//
//	log.Printf("ActionSelect text:%v\n",text)
//	if text=="今日榜单"{
//		miraiHttp.SendText("今日份r18榜单是吧，等会就来，你先把裤子脱好吧")
//		BotService.SendDailyPic()
//		return
//	}
//	if text=="今日女性榜单"{
//		miraiHttp.SendText("今日份女性榜单是吧，等会就来，你先把裤子脱好吧")
//		BotService.SendDailyPic2()
//		return
//	}
//	if text=="知乎热榜"{
//		miraiHttp.SendText("你怎么不看瑟图看这些，我知道了，是贤者模式！")
//		BotService.ZhiHuSend()
//		return
//	}
//	if text=="微博热搜"{
//		miraiHttp.SendText("你会觉得厕所里面会有什么好东西？姥爷你真是独具慧眼啊！")
//		BotService.WeiBoSend()
//		return
//	}
//	if text=="link start"{
//		//miraiHttp.SendText("现在的功能还很少呢~大伙快想想还有什么好idea呀！")
//		caidan:="1.来 点/份 XX 色/瑟图：XX为关键字，如果没有，则随机瑟图（别看那么多，小心群主被抓走！）\\n" +
//			"2.今日榜单:p站r18日榜，只截取前5张.\\n" +
//			"3.知乎热榜：懂的都懂。\\n" +
//			"4.微博热搜：没啥好说的。"
//		url:="https://gitee.com/liang_zi_hao1997/FacemaskOrder/blob/master/weiboLogo.png"
//		miraiHttp.SendMix(caidan,url)
//		return
//
//	}
//
//
//
//	if len(key)==0{
//		return nil
//	}
//	if len(key[0])==0{
//		//没有识别出
//		return nil
//	}
//	if len(key[1])==0{
//		BotService.SendPic()
//	}else{
//
//		log.Printf("ActionSelectKey1:%v\n",string(key[1]))
//		BotService.SendPicWithKey(string(key[1]))
//	}
//	return errors.New("出错啦!")
//}


func ActionSelectWithQQGroup(text string,qqGroup string)(err error)  {
	isWhite:=false
	//qqSendGroups:=[]string{"1085171553","763091038"}
	for _, v := range config.Myconfig.TargetGroups {
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

	//if text=="bcr"{
	//	//miraiHttp.SendTextByQQ("奖励…不能错过…",qqGroup)
	//	miraiHttp.SendPicByQQ("",qqGroup)
	//	return
	//}
	if text=="今周榜单"{
		miraiHttp.SendTextByQQ("奖励…不能错过…",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"weekly")
		return
	}
	if text=="今日榜单"{
		miraiHttp.SendTextByQQ("每日任务…不要忘记…",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"daily")
		return
	}
	if text=="今月榜单"{
		miraiHttp.SendTextByQQ("…要去看看吗？",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"monthly")
		return
	}
	if text=="今周r18榜单"{
		miraiHttp.SendTextByQQ("欢迎回来，哥哥……今天也要开……那个词怎么说来着……开冲了吗？",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"weekly_r18")
		return
	}
	if text=="今日r18榜单"{
		miraiHttp.SendTextByQQ("后方支援就交给我吧…阿比…会加油的…",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"daily_r18")
		return
	}
	if text=="今日男性r18榜单"{
		miraiHttp.SendTextByQQ("大无语事件",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"male_r18")
		return
	}
	if text=="今日女性r18榜单"{
		miraiHttp.SendTextByQQ("北斋：不要给阿比画这些辣！对我来说还太早了！",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"female_r18")
		return
	}
	if text=="今日奇怪r18榜单"{
		miraiHttp.SendTextByQQ("哥哥这么做，一定有其中的深意吧",qqGroup)
		BotService.SendDailyPicWithQQGroupByMode(qqGroup,"r18g")
		return
	}
	if text=="更新配置"{
		config.Myconfig=*config.Myconfig.GetConf()
		BotService.SetBikaToken()
		miraiHttp.SendTextByQQ("哥哥这么做，一定有其中的深意吧",qqGroup)
		return
	}

	if text=="福利"{
		rand.Seed(time.Now().UnixNano())
		take:=5+rand.Intn(5)
		for i:=0;i<take;i++{
			BotService.SendCosWithQQGroup(qqGroup)
		}
		return
	}


	if text=="知乎热榜"{
		miraiHttp.SendTextByQQ("哥哥辛苦了…膝枕…也可以的喔？",qqGroup)
		BotService.ZhiHuSendByQQ(qqGroup)
		return
	}
	if text=="微博热搜"{
		miraiHttp.SendTextByQQ("阿比…这次有帮上哥哥的忙吗？",qqGroup)
		BotService.WeiBoSendByQQ(qqGroup)
		return
	}
	if text=="开启定时色图"{
		config.GroupSets[qqGroup].IsOnTime=true
		miraiHttp.SendTextByQQ("…好可怕…阿比可以回家吗 …哥哥？",qqGroup)
		return
	}
	if text=="关闭定时色图"{
		config.GroupSets[qqGroup].IsOnTime=false
		miraiHttp.SendTextByQQ("unikon已经在这里等候欧尼酱多时了",qqGroup)
		return
	}
	if text=="重置主题"{
		config.GroupSets[qqGroup].Tittle="随机"
		miraiHttp.SendTextByQQ("阿比恢复出厂设置成功！！",qqGroup)
		return
	}
	if text=="查看主题"{
		tittle:=config.GroupSets[qqGroup].Tittle
		miraiHttp.SendTextByQQ("依我所知，北斋现在研究的xp系统是："+tittle,qqGroup)
		return
	}

	if text=="订阅福利姬"{
		config.GroupSets[qqGroup].IsWeiBoFuLiJi=true
		miraiHttp.SendTextByQQ("订阅成功！！",qqGroup)
		return
	}
	if text=="退订福利姬"{
		config.GroupSets[qqGroup].IsWeiBoFuLiJi=false
		miraiHttp.SendTextByQQ("退订成功！",qqGroup)
		return
	}
	if text=="订阅色图"{
		config.GroupSets[qqGroup].IsWeiBoSeTu=true
		miraiHttp.SendTextByQQ("订阅成功！！",qqGroup)
		return
	}
	if text=="退订色图"{
		config.GroupSets[qqGroup].IsWeiBoSeTu=false
		miraiHttp.SendTextByQQ("退订成功！",qqGroup)
		return
	}
	if text=="订阅沙雕图"{
		config.GroupSets[qqGroup].IsWeiBoShaDiaoTu=true
		miraiHttp.SendTextByQQ("订阅成功！！",qqGroup)
		return
	}
	if text=="退订沙雕图"{
		config.GroupSets[qqGroup].IsWeiBoShaDiaoTu=false
		miraiHttp.SendTextByQQ("退订成功！",qqGroup)
		return
	}
	if text=="喝水提醒"{
		config.GroupSets[qqGroup].IsWater=true
		miraiHttp.SendTextByQQ("喝水提醒成功！！",qqGroup)
		return
	}
	if text=="我喝吐了"{
		config.GroupSets[qqGroup].IsWater=false
		miraiHttp.SendTextByQQ("就算不要我提醒，欧尼酱也要多喝水呀！",qqGroup)
		return
	}
	if text=="阿比"||text=="菜单"{
		//miraiHttp.SendText("现在的功能还很少呢~大伙快想想还有什么好idea呀！")
		caidan:=
			"欧尼酱…想要更多了解阿比吗？\\n"+
			"阿比现在装载了以下功能，希望能满足欧尼酱\\n" +
			"1.点图功能：“来点jk瑟图” “来点瑟图” “来点明日方舟瑟图” 可发挥想象力找想看的\\n" +
			"2.榜单功能(维护中):\\n" +
			"      “今日/周/月榜单”\\n" +
			"      ”今日/周r18榜单“\\n" +
			"      “今日男/女性r18榜单”\\n" +
			"3.新闻相关：“知乎热榜” “微博热搜”\\n" +
			"4.定时瑟图：“开启定时色图” “关闭定时色图” “设置白丝主题” “重置主题” “查看主题” ”设置黑丝主题“等。\\n"+
			"5.配菜功能：“查看主题” “设置黑丝主题” “上色图” 根据主题推送若干张图 。\\n"+
			"6.订阅功能：“订阅福利姬/色图/沙雕图” “退订福利姬/色图/沙雕图” \\n"+
			"7.喝水功能：“喝水提醒” “我喝吐了” \\n"+
			"8.福利姬功能：“福利” （三次元给爷爬） \\n"

		//url:="https://gitee.com/liang_zi_hao1997/FacemaskOrder/blob/master/weiboLogo.png"
		miraiHttp.SendMixByQQ(caidan,qqGroup)
		return

	}
	if text=="色图time"||text=="上色图" {
		num:=utils.GetRand(3,7)
		numString:=strconv.Itoa(num)
		miraiHttp.SendTextByQQ("我跟北斋说，我喜欢硬邦邦的欧尼酱，然后它就给我画了这"+numString+"张图，说要递给欧尼酱。",qqGroup)
		if config.GroupSets[qqGroup].Tittle==""{
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
		sendText:="北斋已经变成"+string(keyTittle[1])+"的形状了!"
		miraiHttp.SendTextByQQ(sendText,qqGroup)

							//if string(keyTittle[1])=="诺艾尔"{
							//	sendText="(///_///)主人，请不要开这种玩笑！（传来脱衣解带的声音）"
							//	miraiHttp.SendTextByQQ(sendText,qqGroup)
							//	sendText="(#。﹏。#)请对诺艾尔温柔点。"
							//
							//}
							//
							//if string(keyTittle[1])=="女仆"{
							//	sendText="( >ρ < ”)你的眼前不就有一个女仆吗！"
							//}
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

package cronTask

import (
	"chipsBot/BotService"
	"chipsBot/config"
	"chipsBot/miraiHttp"
	"fmt"
	"github.com/robfig/cron/v3"
)


var CronTask *cron.Cron

func InitCronTask()*cron.Cron{
	fmt.Println("初始化定时任务")
	c := cron.New(cron.WithSeconds()) //精确到秒
	//定时任务
	//spec := "* 1 * * * ?"

	//c.AddFunc(spec, func() {
	//	fmt.Println("开始执行更新临时文件为待回收状态定时任务")
	//	err:=service.RecycleFiles()
	//	if err!=nil{
	//		fmt.Println("执行更新临时文件为待回收状态定时任务出错")
	//	}
	//})
	//c.AddFunc(spec, func() {
	//	fmt.Println("开始执行待删除文件定时任务")
	//	err:=service.DeleteFiles()
	//	if err!=nil{
	//		fmt.Println("开始执行待删除文件定时任务出错")
	//	}
	//})
	//c.AddFunc("*/5 * * * * ?", func() {
	//	fmt.Println("每五秒心跳一次")
	//})

	c.AddFunc("0 45 */1 * * ?", func() {//cron表达式，每15min一次
		fmt.Println("开始执行发图定时任务")
		for _, v := range config.QQSendGroups {
			if config.GroupSets[v].IsOnTime==true{
				if config.GroupSets[v].Tittle!="随机"{
					for i := 0; i <= 5; i++{BotService.SendPicWithKeyAndQQGroup(config.GroupSets[v].Tittle,v)}
				}else{
					for i := 0; i <= 5; i++{BotService.SendPicByQQ(v)}
				}
			}
		}

	})

	c.AddFunc("0 0 15 * * ?", func() {//
		for _, v := range config.QQSendGroups {
			text:= "主人，三点多啦！来喝一杯靓靓的茶，老板不会疼你，但诺艾尔会！"

			miraiHttp.SendTextByQQ(text,v)
			text="在这个时间喝茶，对人体能起到调理的作用，增强身体的抵抗力、还能防止感冒，此时喝茶是一天中最重要的，俗称下午茶，" +
				"听说这在主人的故乡叫申时茶。对一些“三高”人群来说，如果坚持喝下午茶，能起到药物都无法达到的效果哦！"
			miraiHttp.SendTextByQQ(text,v)
		}
	})
	c.AddFunc("0 30 20 * * ?", func() {//
		for _, v := range config.QQSendGroups {
			text:= "主人，别忘了要定时喝水哦！"
			miraiHttp.SendTextByQQ(text,v)
			text="这个时间是人体免疫系统最活跃的时间，如果能喝上一泡茶，人体会很容易修补和恢复免疫系统，再造细胞等。但千万不要喝绿茶或清香型铁观音，因为绿茶是不发酵茶，对人体有一定的刺激，清香型铁观音性近绿茶。"
			miraiHttp.SendTextByQQ(text,v)
			text="还有未发酵的白茶、生普洱都不应该喝。应该选择喝黑茶或熟普洱、陈年铁观音、老白茶、六堡茶、茯砖。这些茶叶性致温纯，不会影响人体正常的睡眠。而且晚饭之后喝全发酵类的茶可帮助分解积聚的脂肪，既暖胃又帮助消化。"
			miraiHttp.SendTextByQQ(text,v)
		}
	})
	c.AddFunc("0 0 9 * * ?", func() {//
		for _, v := range config.QQSendGroups {
			text:= "主人，诺艾尔给你斟茶了，经过了一整个晚上的休息之后，想必您消耗了大量的水分，血液的浓度变大，喝上一杯淡茶水不但能快速的补充身体所需的水分，清理肠胃，还可以降低血压，稀释血液，有益健康"
			miraiHttp.SendTextByQQ(text,v)
		}
	})

	c.AddFunc("0 0 18 * * ?", func() {//
		for _, v := range config.QQSendGroups {
			text:= "傍晚了，主人是要先吃饭，还是先洗澡？"
			miraiHttp.SendTextByQQ(text,v)
		}
	})
	c.AddFunc("0 0 0 * * ?", func() {//
		for _, v := range config.QQSendGroups {
			text:= "夜深了，主人也该睡了，放心吧，诺艾尔也会在梦里守候主人的哦~"
			miraiHttp.SendTextByQQ(text,v)
		}
	})

	//c.AddFunc("0 30 */8 * * ?", func() {//cron表达式，每8小时一次
	//
	//	fmt.Println("开始执行热搜定时任务")
	//	BotService.ScienceSend()
	//})
	//c.AddFunc("0 0 */10 * * ?", func() {//cron表达式，每10小时一次
	//
	//	fmt.Println("开始执行科学探索定时任务")
	//	BotService.ScienceSend()
	//})
	c.Start()
	return c
}


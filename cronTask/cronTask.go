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
		for _, v := range config.Myconfig.TargetGroups {
			if config.GroupSets[v].IsOnTime==true{
				if config.GroupSets[v].Tittle!="随机"{
					for i := 0; i <= 5; i++{BotService.SendPicWithKeyAndQQGroup(config.GroupSets[v].Tittle,v)}
				}else{
					for i := 0; i <= 5; i++{BotService.SendPicByQQ(v)}
				}
			}
		}

	})
	c.AddFunc("0 0/20 * * * ?", func() {//cron表达式，每15min一次
		fmt.Println("开始执行订阅任务")
		for _, qq:= range config.Myconfig.TargetGroups {
			var tempQQgroups []string
			if config.GroupSets[qq].IsWeiBoSeTu==true{
				tempQQgroups=append(tempQQgroups,qq)
			}
			BotService.SendBlogsByUserIdsByQQGroups(config.Myconfig.WeiBoSeTuUps,tempQQgroups)
		}
	})
	c.AddFunc("0 7/20 * * * ?", func() {//cron表达式，每15min一次

		for _, qq:= range config.Myconfig.TargetGroups {
			var tempQQgroups []string
			if config.GroupSets[qq].IsWeiBoShaDiaoTU==true{
				tempQQgroups=append(tempQQgroups,qq)
			}
			BotService.SendBlogsByUserIdsByQQGroups(config.Myconfig.WeiBoShaDiaoUps,tempQQgroups)
		}

	})
	c.AddFunc("0 14/20 * * * ?", func() {//cron表达式，每15min一次
		fmt.Println("开始执行订阅任务")
		for _, qq:= range config.Myconfig.TargetGroups {
			var tempQQgroups []string
			if config.GroupSets[qq].IsWeiBoFuLiJi==true{
				tempQQgroups=append(tempQQgroups,qq)
			}
			BotService.SendBlogsByUserIdsByQQGroups(config.Myconfig.WeiBoFuLiJiUps,tempQQgroups)
		}

	})



	c.AddFunc("0 0 15 * * ?", func() {//
		for _, v := range config.Myconfig.TargetGroups {
			text:= BotService.DrinkTeaContent()
			if config.GroupSets[v].IsWater{
				miraiHttp.SendTextByQQ(text,v)
			}
		}
	})
	c.AddFunc("0 0 12 * * ?", func() {//
		for _, v := range config.Myconfig.TargetGroups {
			text:= "欧尼酱，就算再忙，午餐也要吃顿好的~如果有什么事，只要喊我的名字就可以了，我会随叫随到。"
			if config.GroupSets[v].IsWater{
				miraiHttp.SendTextByQQ(text,v)
			}
		}
	})
	c.AddFunc("0 30 7-23 * * ?", func() {//
		for _, v := range config.Myconfig.TargetGroups {
			text:= BotService.WaterContent()
			if config.GroupSets[v].IsWater{
				miraiHttp.SendTextByQQ(text,v)
			}
		}
	})

	c.AddFunc("0 0 0 * * ?", func() {//
		config.Myconfig=*config.Myconfig.GetConf()
	})


	c.AddFunc("0 0 18 * * ?", func() {//
		for _, v := range config.Myconfig.TargetGroups {
			text:= "傍晚了，独角兽想和欧尼酱商量一下今晚吃什么"
			if config.GroupSets[v].IsWater{
				miraiHttp.SendTextByQQ(text,v)
			}

		}
	})
	c.AddFunc("0 0 0 * * ?", func() {//
		for _, v := range config.Myconfig.TargetGroups {

			text:= "夜深了，欧尼酱也会在梦里陪着独角兽的吧~"
			if config.GroupSets[v].IsWater{
				miraiHttp.SendTextByQQ(text,v)
			}
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


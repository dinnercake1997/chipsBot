package cronTask

import (
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

	//c.AddFunc("0 */60 * * * ?", func() {//cron表达式，每15min一次
	//
	//	fmt.Println("开始执行发图定时任务")
	//	for i := 0; i <= 5; i++{
	//		BotService.SendPic()
	//	}
	//
	//})
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


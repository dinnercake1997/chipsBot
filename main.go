package main

import (
	"chipsBot/cronTask"
	"fmt"
	"net/http"
)

func main(){

	//for{
	//	utils.SendPic()
	//	log.Printf("每5分钟打印一次\n")
	//	time.Sleep(time.Second*60*15)
	//}
	cronTask.CronTask=cronTask.InitCronTask()
	http.HandleFunc("/", IndexHandler)
	err:=http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		// 服务器创建失败
		panic("服务器创建失败")
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}


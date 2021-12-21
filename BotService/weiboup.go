package BotService

import (
	"chipsBot/config"
	"chipsBot/miraiHttp"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)
type NewWeiBoInfo struct{
	userName string
	blogs []Blog

}
type Blog struct{
	text string
	originPic string
	pics []string
}

func  SendBlogsByUserIdsByQQGroups(userIDs []string,qqGroups []string ) (err error) {
	log.Printf("总共有%v个微博用户需要搜索!\n", len(userIDs))
	for index,userId:=range userIDs{
		log.Printf("第%v个userid:%s!\n",index,userId)
		containerId,err:=GetWeiboContainId(userId)
		if err!=nil{
			log.Println("GetWeiboContainIdErr:",err)
			return err
		}
		newWeiBoInfos,err:=GetNewBlogsByUserIdAndContainid(userId,containerId)
		if err!=nil{
			log.Println("GetNewBlogsByUserIdAndContainidErr:",err)
			return err
		}
		if len(newWeiBoInfos.blogs)!=0{
			for i:=0;i<len(newWeiBoInfos.blogs);i++{
				pics:=newWeiBoInfos.blogs[i].pics
				for j:=0;j<len(pics);j++{
					picUrl:=pics[j]
					for _,qqGroup:=range qqGroups{
						miraiHttp.SendPicByQQ(picUrl,qqGroup)
					}

				}
			}
		}
	}
	return err
}

func GetWeiboContainId(weiBoUserId string)  (containerId string,err error){

	url := "https://m.weibo.cn/api/container/getIndex?type=uid&value="+weiBoUserId

	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	if err!=nil{
		log.Printf(err.Error())
		log.Printf("请求微博用户container出错1111")
		err=errors.New("请求微博用户container出错111")
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	json,err:= simplejson.NewJson(body)
	if err!=nil{
		log.Printf("请求微博用户container出错转json出错:%v",err)
		return "",errors.New("请求微博用户container出错转json出错")
	}
	if err!=nil{
		return
	}
	//groupqq,_:=json.Get("sender").Get("group").Map()
	////groupQQ=groupQQInt.(string)
	//groupQQ=fmt.Sprintf("%v",groupqq["id"])
	userInfo,_:=json.Map()
	okGetUserInfo:=fmt.Sprintf("%v",userInfo["ok"])
	fmt.Printf("okGetUserInfo:%s\n",okGetUserInfo)

	if okGetUserInfo!="1" {
		return "",errors.New("请求微博用户container出错")
	}
	blogContainJsonMap, _ :=json.Get("data").Get("tabsInfo").Get("tabs").Array()
	for _,item :=range blogContainJsonMap {
		dataMap,ok:= item.(map[string]interface{})
		if !ok{
			log.Printf("请求loliconAPIkey获取图片url出错:%v",err)
			return "",errors.New("请求求微博热点接口获取图片信息出错")
		}
		id:=fmt.Sprintf("%v",dataMap["id"])
		containerId,_=dataMap["containerid"].(string)
		log.Printf("id:%v",id)
		log.Printf("containerid:%s",containerId)
		if id=="2" {
			return containerId,nil
		}

	}
	fmt.Println("获取用户最新微博containerId失败")
	return "",errors.New("获取用户最新微博containerId失败！")
}

func GetNewBlogsByUserIdAndContainid(userId string,containerid string) (newWeiBoInfo NewWeiBoInfo ,err error) {
	var blogs []Blog
	url := "https://m.weibo.cn/api/container/getIndex?type=uid&value="+userId+"&containerid="+containerid
	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err!=nil{
		log.Printf("请求用户微博接口出错")
		err=errors.New("请求用户微博接口出错")
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	//fmt.Printf("body:%s",string(body))
	jsonT,err:= simplejson.NewJson(body)
	if err!=nil{
		log.Printf("请求用户微博接口出错转json出错:%v",err)
		return newWeiBoInfo,errors.New("请求用户微博接口出错转json出错")
	}
	if err!=nil{return}
	//groupqq,_:=json.Get("sender").Get("group").Map()
	////groupQQ=groupQQInt.(string)
	//groupQQ=fmt.Sprintf("%v",groupqq["id"])
	blogsMap,_:=jsonT.Map()
	blogsInfo:=fmt.Sprintf("%v",blogsMap["ok"])
	fmt.Printf("blogsInfo:%s\n",blogsInfo)
	if blogsInfo!="1" {
		if blogsInfo=="0"{
			return newWeiBoInfo,nil
		}
		return newWeiBoInfo,errors.New("请求微博用户container出错")
	}

	cardsArrary, _ :=jsonT.Get("data").Get("cards").Array()
	for _,card:=range cardsArrary{
		cardMap,_:= card.(map[string]interface{})
		dataType , _ := json.Marshal(cardMap)
		dataString := string(dataType)
		//fmt.Println("dataString:",dataString)
		//mblogJsonString:=fmt.Sprintf("%v",cardMap["mblog"].(string))
		//fmt.Printf("mblogJsonString:%v\n",mblogJsonString)
		mblogJson,_:= simplejson.NewJson([]byte(dataString))
		//mblogJson, _ :=json.Marshal(mblogJsonString)
		//fmt.Printf("mblogJson:%v\n",mblogJson)
		mblogMap,_:=mblogJson.Get("mblog").Map()
		tempBlog:=new(Blog)
		isSend:=false
		if mblogMap["created_at"]!=nil{
			isSend=CheckIsNewBlog(mblogMap["created_at"].(string))
		}
		if isSend{
			tempBlog.originPic,_=mblogJson.Get("mblog").Get("original_pic").String()
			tempBlog.pics,_=mblogJson.Get("mblog").Get("pic_ids").StringArray()
			if len(tempBlog.pics)!=0{
				tempPic:=tempBlog.pics[0]
				for index,pic:=range tempBlog.pics{
					tempBlog.pics[index]=strings.Replace(tempBlog.originPic, tempPic, pic, 1)
				}
			}
			tempBlog.text=mblogMap["text"].(string)
			blogs=append(blogs,*tempBlog)
			fmt.Printf("text:%v\n",tempBlog.text)
			for _,pic :=range  tempBlog.pics{
				fmt.Printf("pic:%v\n",pic)
			}
		}
		newWeiBoInfo.userName=""
		newWeiBoInfo.blogs=blogs
	}
	return


}

func CheckIsNewBlog(timeString string )( res bool ){
	//tempArrary:=strings.Split(timeString,":")
	//numBlog,_:= strconv.Atoi(tempArrary[1])
	//timeNow:=time.Unix(time.Now().Unix(),0).Format("2006-01-02 15:04:05")//模板要用go诞生的日期
	//numNow,_:=strconv.Atoi(timeNow)
	isNewSeconds:= config.Myconfig.IsNewWeiBoSeconds
	timeStampBlog:=WeiBoTimeToTimeStamp(timeString)
	timeNow:= time.Now().Unix()
	//fmt.Printf("timeStampNow:%v\n",timeNow)
	//fmt.Printf("timeStampBlog:%v\n",timeStampBlog)
	temp:=timeNow-timeStampBlog
	fmt.Printf("res:%v\n",temp)
	if temp>isNewSeconds{
		res=false
	}else {
		res=true
		log.Printf("新微博：%v\n",temp)
	}
	return res
}

func WeiBoTimeToTimeStamp(timeString string)(timeStamp int64){
	MonthMap:= map[string]string{}
	MonthMap["Jan"]="01"
	MonthMap["Feb"]="02"
	MonthMap["Mar"]="03"
	MonthMap["Apr"]="04"
	MonthMap["May"]="05"
	MonthMap["Jun"]="06"
	MonthMap["Jul"]="07"
	MonthMap["Aug"]="08"
	MonthMap["Sep"]="09"
	MonthMap["Oct"]="10"
	MonthMap["Nov"]="11"
	MonthMap["Dec"]="12"
	tempArrary1:=strings.Split(timeString," ")
	//tempArrary2:=strings.Split(timeString,":")

	month:= MonthMap[tempArrary1[1]]
	//log.Printf("month0:%v",tempArrary1[1])
	//log.Printf("month:%v",month)
	day:=tempArrary1[2]
	timeBlog:=tempArrary1[3]
	year:=tempArrary1[5]

	timeStringNormal := year + "-" + month + "-" + day + " " + timeBlog
	log.Printf("time:%S",timeStringNormal)
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, timeStringNormal, loc) //使用模板在对应时区转化为time.time类型
	timeStamp = theTime.Unix()                                          //转化为时间戳 类型是int64
	//fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	//fmt.Println(timeStamp)                                                 //打印输出时间戳 1420041600
	//log.Printf("timeStamp:%S",timeStamp)
	return
}
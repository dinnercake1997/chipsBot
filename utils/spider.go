package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//func SpiderTest(){
//	resp,err:=http.Get("http://www.baidu.com/")
//	if err!=nil{
//		fmt.Println("http get err",err)
//		return
//	}
//	body,err :=ioutil.ReadAll(resp.Body)
//	if err!=nil{
//		fmt.Println("ioutil ReadAll err",err)
//		return
//	}
//	fmt.Println(string(body))
//}
//
//func Download(url string)(bodyString string,err error){
//	client := &http.Client{}
//	req, err := http.NewRequest("GET", url, nil)
//	// 自定义Header
//	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
//
//	resp, err := client.Do(req)
//	if err != nil {
//		fmt.Println("http get error", err)
//		return
//	}
//	//函数结束后关闭相关链接
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println("read error", err)
//		return
//	}
//	fmt.Println(string(body))
//	bodyString=string(body)
//	return
//}
//
//
func TestMeiNv(){
	client := &http.Client{}
	begin := time.Now()
	//for page := 1; page <= 10; page++ {
		url := "https://www.tuaoo.cc/"
		request, err := http.NewRequest("GET", url, nil)
		request.Header.Add("User-Agent", "okhttp/3.8.1")
		//request.Header.Add("User-Agent", "okhttp/3.8.1")
		request.Header.Add("Content-Type", "application/json; charset=UTF-8")
		response, err := client.Do(request)
		fmt.Printf("response.Body:%v\n",response)
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("http read error.")
		}
		src := string(body)
		fmt.Println(src)
			opending, err := goquery.NewDocumentFromResponse(response)
			if err != nil {
				log.Fatal(err)
			}
		ele := opending.Find(".imgc")
		ele.Each(func(index int, content *goquery.Selection) {
			//name, _ := content.Find("a").First().Attr("title")
			fmt.Printf("1")
		})
	//}
	end := time.Now()
	spendTime := end.Sub(begin)
	fmt.Println("花费时间为:", spendTime)

}
//func TestQ(){
//	client := &http.Client{}
//	begin := time.Now()
//
//		url := "https://www.qiushibaike.com/hot/"
//		request, err := http.NewRequest("GET", url, nil)
//		request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36")
//		request.Header.Add("Host", "www.qiushibaike.com")
//		response, err := client.Do(request)
//
//		//fmt.Printf("response.Body:%v\n",response.Body)
//		opending, err := goquery.NewDocumentFromResponse(response)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Printf("opending:%v\n",opending)
//		ele := opending.Find("h2")
//		fmt.Printf("len:%v\n",ele.Length())
//		ele.Each(func(index int, content *goquery.Selection) {
//			name:= content.Text()
//			fmt.Printf("%d: %s\n", index, name)
//		})
//
//	end := time.Now()
//	spendTime := end.Sub(begin)
//	fmt.Println("花费时间为:", spendTime)
//
//}
//
//func TestWeiBo(){
//	client := &http.Client{}
//	begin := time.Now()
//
//	url := "https://m.weibo.cn/u/"+config.Myconfig.WeiBoUps[0]
//	fmt.Printf("url:%s\n",url)
//	request, err := http.NewRequest("GET", url, nil)
//	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36")
//	//request.Header.Add("Host", "m.weibo.cn")
//	request.Header.Add("referer", "https://m.weibo.cn/")
//	request.Header.Add("Cookie", "https://m.weibo.cn/")
//	response, err := client.Do(request)
//
//	body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		fmt.Println("read error", err)
//		return
//	}
//	fmt.Println("body:")
//	fmt.Println(string(body))
//
//	opending, err := goquery.NewDocumentFromResponse(response)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("opending:%v\n",opending)
//
//	//class="card m-panel card9 weibo-member card-vip"
//
//	ele := opending.Find("span")
//	fmt.Printf("len:%v\n",ele.Length())
//	ele.Each(func(index int, content *goquery.Selection) {
//		//time:= content.Find(".time").Text()
//		time:= content.Text()
//		fmt.Printf("%d: %s\n", index, time)
//	})
//
//	end := time.Now()
//	spendTime := end.Sub(begin)
//	fmt.Println("花费时间为:", spendTime)
//
//
//}


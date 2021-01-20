package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"

	"log"
	"mime/multipart"
	"net/http"
)

func DoGet(url string)(content string ,err error){

	req, _ := http.NewRequest("GET", url, nil)
	//请求头添加accessToken
	//req.Header.Add("Cookie", token)
	res, err := (&http.Client{}).Do(req)
	//调用htto请求出错
	if err != nil {
		err=errors.New("CheckAccessToken调用htto请求出错！")
		return "",err

	}
	robots, err := ioutil.ReadAll(res.Body)
	if err!=nil {
		err=errors.New("读取核心返回信息失败！")
		return "",err
	}
	log.Println("请求返回信息：",string(robots))
	resultMap := make(map[string]string)
	//返回信息转map
	err= json.Unmarshal(robots, &resultMap)
	//查看返回信息的map对象是否有rows字段，有表示accessToken有效
	_, ok := resultMap["code"]
	res.Body.Close()
	if ok{
		return string(robots),nil
	}else{
		err=errors.New("失败！")
		return
	}

}
func DoPost(url string,requestBody  string)(content string ,err error){
	log.Println("开始请求cqhttp")
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	log.Printf("requestBody：%s",requestBody)
	jsonStr := []byte(requestBody)
	err= writer.Close()
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("请求出错")
		return "请求出错", err
	}
	//token:="access_token="+accessToken
	//req.Header.Add("Cookie", token)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("请求出错")
		return "请求出错", err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	content=string(bytes)
	log.Println(content)
	if err != nil {
		log.Println("读取响应体出错")
		return "读取响应体出错", err
	}
	resp.Body.Close()
	return content, nil

}
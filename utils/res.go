
package utils

import (
	"encoding/json"
)

type ResMsgAndCode struct {
	Msg   string`json:"Msg"`
	Code int`json:"Code"`
	Text string  `json:"Text"`
}

//result=0表示请求成功但是业务失败，code=0表示请求失败
func SetResMsgAndCode(msg string,code int,text string )(resJson[]byte){
	res:=new(ResMsgAndCode)
	res.Msg=msg
	res.Code=code
	res.Text=text
	resJson,_=json.Marshal(res)
	return resJson
}


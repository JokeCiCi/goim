package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type H struct {
	Code int         `json:"code"` // json序列化时键小写
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"` // omitempty,nil不序列化为json
}

// 封装返回json
func Resp(w http.ResponseWriter, code int, msg string, data interface{}) {
	h := &H{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	b, err := json.Marshal(h) // 结构体转json
	if err != nil {
		log.Println(err.Error())
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json") // 设置响应头，数据json格式
	w.Write(b)                                         // 输出
}

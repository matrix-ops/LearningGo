package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

var post string = "{\"pageNo\":1,\"pageSize\":20}"
var url string = "https://onlineposuat.taikoohui.com/taikoomall-member-app/parking/list"

func main() {
	benchmark(url)
}
func benchmark(url string) {
	for {
		postJson := []byte(post)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(postJson))
		// NewRequest 的第三个参数要求是一个io.Reader类型的
		// 这个类型是一个接口类型，任何类型只要实现了Read方法都能实现这个接口
		// Reader接口类型的定义如下：
		// type Reader interface {
		// 		Read(p []byte) (n int, err error)
		// }
		// n是总共读取的数据量
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{
			Timeout: time.Second * 3,
		}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(resp)
			fmt.Println("Requests Success")
		}
	}
}

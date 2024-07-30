package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// 定义一个结构体来表示接收到的JSON数据
type InputData struct {
	// 包含用户发送的其他字段，可以是任意字段，这里只作为示例
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 定义一个结构体来表示返回的JSON数据
type OutputData struct {
	InputData
	Time int64 `json:"time"`
}

// 处理函数
func handler(w http.ResponseWriter, r *http.Request) {
	// 设置Content-Type为application/json
	w.Header().Set("Content-Type", "application/json")

	// 解析JSON请求体
	var inputData InputData
	err := json.NewDecoder(r.Body).Decode(&inputData)
	// 这是将接收到的json请求解码到一个结构体里面
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 创建返回的数据结构体，添加当前时间戳
	outputData := OutputData{
		InputData: inputData,
		Time:      time.Now().Unix(),
	}

	// 将结构体编码成JSON并返回
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(outputData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// 路由和处理函数
	http.HandleFunc("/processjson", handler)

	// 启动服务器
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

var requestCounter uint64

func sendRequest(wg *sync.WaitGroup, client *http.Client, reqNum uint64) {
	defer wg.Done()
	// 将请求编号添加为查询参数
	url := fmt.Sprintf("https://openapi.payeco.com/receipt-app-demo/demo/test/test?request_number=%d", reqNum)
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("请求 #%d 出错: %v\n", reqNum, err)
		return
	}
	defer resp.Body.Close()

	log.Printf("请求 #%d 响应状态: %s\n", reqNum, resp.Status)
}

func main() {
	// 打开日志文件
	logFile, err := os.OpenFile("request_logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("无法打开日志文件: %v", err)
	}
	defer logFile.Close()

	// 设置日志输出到文件
	log.SetOutput(logFile)

	client := &http.Client{}
	var wg sync.WaitGroup
	ticker := time.NewTicker(1 * time.Second)

	for range ticker.C {
		wg.Add(50)
		for i := 0; i < 50; i++ {
			// 使用原子操作安全地递增请求计数器
			reqNum := atomic.AddUint64(&requestCounter, 1)
			go sendRequest(&wg, client, reqNum)
		}
		wg.Wait()
	}
}

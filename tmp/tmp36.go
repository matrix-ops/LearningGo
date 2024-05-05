package main

import "fmt"

func main() {
	c1 := make(chan int)
	go func() {
		c1 <- 1
		close(c1)
		// 关闭c1
		// 关闭的好处是接收方不用使用额外的手段来确认c1是否还有数据没发送
		fmt.Println("c1 is closed")
	}()
	fmt.Println(<-c1)
}

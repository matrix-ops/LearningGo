package main

import (
	"fmt"
	"time"
)

func goRoutineTest(v int, c chan int) {
	time.Sleep(time.Second * 3)
	c <- v
}
func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go goRoutineTest(i, c)
	}
	// 要摆脱线性思维去理解goroutine，在这里
	// 是同时启动了10个异步的goroutine往main goroutine发送通道的值
	// 因为通道c是非缓冲的通道，因此一次只能接受一个值
	// 在休眠了3秒钟之后，10个goroutine在不停地轮询尝试将v发送到c
	// 之所以看上去好像是一瞬间全部的值都被打印出了出来，是因为代码太简单，输出太快的原因，根本不需要什么计算
	// 下面的循环是一次一个接受通道上的值
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println(<-c)
	}
}

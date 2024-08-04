package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg_8 sync.WaitGroup

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}
func main() {
	wg.Add(2)
	b := make(chan int)
	go Player("zhangweilong1", b)
	go Player("zhangweilong2", b)
	b <- 1
	wg_8.Wait()
}

func Player(name string, ball chan int) {
	defer wg_8.Done()
	var winOrLose int
	for {
		number, ok := <-ball
		if !ok {
			// fmt.Println(name, " Wins", "number is ", number)
			// 如果这样写，number将永远是0，因为从一个已经关闭的通道读取值从永远读到0
			fmt.Println(name, " wins")
			return
		}
		winOrLose = rand.Intn(1000)
		if winOrLose%13 == 0 {
			fmt.Println(name, "Generated number is ", winOrLose, " ", name, " lose ")
			close(ball)
			return
		}
		fmt.Println(name, " Hits: ", number, "number is ", winOrLose)
		number++
		ball <- number
	}
}

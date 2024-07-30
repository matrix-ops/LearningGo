package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}
func main() {
	wg.Add(2)
	count := make(chan int)
	go player("zhangweilong", count)
	go player("wangjinxi", count)
	count <- 1
	wg.Wait()
}

func player(name string, count chan int) {
	defer wg.Done()
	for {
		ball, ok := <-count
		if !ok {
			fmt.Println("Player ", name, "won")
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			close(count)
			fmt.Println("Player ", name, "lose ", "本次是第", ball, "次击球，导致本次击球失败的随机数是：", n)
			return
		}
		fmt.Println(name, "第", ball, "次击球")
		ball++
		count <- ball
	}
}

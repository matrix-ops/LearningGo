package main

import (
	"fmt"
	"sync"
	"time"
)

var wg_9 sync.WaitGroup

func main() {
	wg_9.Add(1)
	baton := make(chan int)
	fmt.Println("现在是第 1 位选手")
	go Runner(baton)
	baton <- 1
	wg_9.Wait()
}

func Runner(baton chan int) {
	runner := <-baton
	var newRunner int
	if runner != 4 {
		newRunner = runner + 1
		fmt.Println("现在是第", newRunner, "位选手")
		go Runner(baton)
	}
	time.Sleep(100 * time.Millisecond)
	if runner == 4 {
		fmt.Println("Game Over")
		wg_9.Done()
		return
	}
	baton <- newRunner
}

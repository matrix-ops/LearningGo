package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	runner int = 0
	wg2    sync.WaitGroup
)

func main() {
	wg2.Add(1)
	baton := make(chan int)
	go Runner2(baton)
	baton <- runner
	wg2.Wait()
}
func Runner2(baton chan int) {
	var newRunner int
	newRunner = <-baton
	newRunner += 1
	if newRunner == 4 {
		fmt.Printf("第 %d 号选手正在赛道上奔跑\n", newRunner)
		wg2.Done()
		return
	} else {
		fmt.Printf("第 %d 号选手正在赛道上奔跑\n", newRunner)
	}
	time.Sleep(100 * time.Millisecond)
	go Runner2(baton)
	if newRunner == 3 {
		baton <- newRunner
		close(baton)
	} else {
		baton <- newRunner
	}
}

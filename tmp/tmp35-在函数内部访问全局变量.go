package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu sync.Mutex
	s1 = 1200
	s2 = 1300
)

func TestFunction1() {
	mu.Lock()
	fmt.Println("Function 1 mu Locked ")
	fmt.Println(s1)
	mu.Unlock()
	fmt.Println("Function 1 mu Unlocked ")
}

func TestFunction2() {
	mu.Lock()
	fmt.Println("Function 2 mu Locked ")
	fmt.Println(s2)
	mu.Unlock()
	fmt.Println("Function 2 mu Unlocked ")
}
func main() {
	for i := 0; i < 5; i++ {
		go TestFunction1()
		go TestFunction2()
	}
	time.Sleep(time.Second * 1)
}

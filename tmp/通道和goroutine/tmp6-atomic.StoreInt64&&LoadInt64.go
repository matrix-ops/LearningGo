package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)
	Dowork("A")
	Dowork("B")
	time.Sleep(time.Second)
	atomic.StoreInt64(&shutdown, 1)
}
func Dowork(name string) {
	defer wg.Done()
	fmt.Println("Doing Work " + name)
	time.Sleep(250 * time.Millisecond)
	if atomic.LoadInt64(&shutdown) == 1 {
		fmt.Println("Work Done " + name)
	}
}

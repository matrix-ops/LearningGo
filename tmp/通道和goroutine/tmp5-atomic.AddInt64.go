package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	atomAdder(12)
	atomAdder(13)
	fmt.Println(counter)
	// 25
}

func atomAdder(id int) {
	defer wg.Done()
	atomic.AddInt64(&counter, int64(id))
}

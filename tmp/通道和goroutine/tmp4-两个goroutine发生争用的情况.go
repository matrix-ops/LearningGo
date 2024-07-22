package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	go intCounter(1)
	go intCounter(2)
	wg.Wait()
	fmt.Println("Counter Value :", counter)
}
func intCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		value := counter
		runtime.Gosched()
		value++
		counter = value
		if counter == 2 {
			fmt.Println(value)
		}
	}
}

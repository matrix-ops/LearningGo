package main

import "fmt"

func returnChan2() <-chan int {
	ch := make(chan int)
	go func() {
		ch <- 10
	}()
	return ch
}
func main() {
	ch := returnChan2()
	fmt.Println(<-ch)
}

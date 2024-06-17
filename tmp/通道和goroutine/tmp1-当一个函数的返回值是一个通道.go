package main

import "fmt"

func returnChan() <-chan int {
	ch := make(chan int)
	go func() {
		ch <- 10
		close(ch)
	}()
	return ch
}
func main() {
	fmt.Println(<-returnChan())
}

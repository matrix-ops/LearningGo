package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 99
	}()
	fmt.Println(<-c)
}

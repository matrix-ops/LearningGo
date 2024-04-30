package main

import "fmt"

func main() {
	var c chan int
	c = make(chan int)
	go func() {
		c <- 99
	}()
	r := <-c
	fmt.Println(r)
}

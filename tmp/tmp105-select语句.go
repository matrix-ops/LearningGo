package main

import (
	"fmt"
	"time"
)

func worker() {
	n := 0
	next := time.After(time.Second * 1)
	for {
		select {
		case <-next:
			n++
			fmt.Println(n)
			next = time.After(time.Second * 1)
		}
	}
}

func main() {
	go worker()
	time.Sleep(time.Second * 100)
}

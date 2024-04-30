package main

import (
	"fmt"
	"time"
)

func sleeper(id int) {
	fmt.Println("I am ", id, "i am sleeping now.")
}
func main() {
	for i := 0; i < 10; i++ {
		go sleeper(i)
	}
	time.Sleep(4 * time.Second)
}

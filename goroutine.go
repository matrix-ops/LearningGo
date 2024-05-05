package main

import (
	"fmt"
	"time"
)

func main() {
	go sleepGopher()
	time.Sleep(4 * time.Second)
}
func sleepGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("... snore ...")
}

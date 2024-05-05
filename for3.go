package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 10; i > 0; i -= 1 {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

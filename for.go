package main

import (
	"fmt"
	"time"
)

func main() {
	var count int = 10
	fmt.Println("倒数10")
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("现在已经倒数了10下")
}

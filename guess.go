package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var randomNumber1 int = 10
	for {
		var randomNumber2 int = rand.Intn(100)
		switch {
		case randomNumber2 > randomNumber1:
			fmt.Println(randomNumber2, "这个数字比我刚刚声明的要大")
		case randomNumber2 < randomNumber1:
			fmt.Println(randomNumber2, "这个数字比我刚刚声明的要小")
		}
		if randomNumber2 == randomNumber1 {
			fmt.Println("randomNumber2", randomNumber2)
			fmt.Println("randomNumber1", randomNumber1)
			fmt.Println("这两个数字相同")
			break
		}
	}
}

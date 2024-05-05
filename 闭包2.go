package main

import (
	"fmt"
	"strings"
)

func main() {
	func(l string) {
		fmt.Println(strings.ToUpper(l))
	}("You son of Bitch")
	// 会以全大写输出
}

package main

import "fmt"

var f = func() {
	fmt.Println("Love you")
}

func main() {
	f()
	print := func(prefix, i string) (o string) {
		fmt.Println(prefix + i)
		return prefix + i
	}
	print("haha ", "love you")
}

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, args := range os.Args[1:] {
		s += args + sep
		sep = " "
	}
	// range一个数组的时候返回两个值，第一个值是索引，第二个值是内容
	fmt.Println(s)
}

package main

import (
	"fmt"
	"os"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i, file := range files {
		// ReadDir返回的第一个参数是一个切片
		// range遍历切片的时候第一个返回值是切片元素的索引
		fmt.Println(i, file)
	}
}

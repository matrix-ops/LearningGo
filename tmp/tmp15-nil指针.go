package main

import (
	"fmt"
)

func main() {
	var t *int
	fmt.Println(t)
	// t里面没有任何内存地址
	// 打印nil
	a := 10
	t = &a
	fmt.Println(t, &a)
	// 会打印变量a的内存地址
	fmt.Println(&t)
	// 打印指针变量t的内存地址
	if t != nil {
		fmt.Println(*t)
	}
}

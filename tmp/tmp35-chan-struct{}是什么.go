package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s struct{}
	c := make(chan struct{}, 1)
	// 这是一个通道，通道的类型是空结构体
	// 空结构体类型可以写成struct{}
	// 空结构体字面量要写成struct{}{}
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(s))
}

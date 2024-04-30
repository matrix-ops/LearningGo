package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "zhangweilong"
	people := &name
	man := *people
	fmt.Println("man的类型是：", reflect.TypeOf(man))
	// string
	// 在通过解引用string指针people并将其值复制给man的过程中，赋值的只是一个字符串
	*people = "liuwenzhuo"
	// 通过解引用，操作people所指向的那片内存地址，将直接操作name的内存地址
	fmt.Println(name)
	// liuwenzhuo
}

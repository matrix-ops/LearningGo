package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name    string
	address string
	age     int
}

func main() {
	timmy := &person{
		name:    "zhangweilong",
		address: "GuangDong GuangZhou",
		age:     19,
	}
	fmt.Println(reflect.TypeOf(timmy))
	// 指针，类型为main.person
	fmt.Println((*timmy).name)
	fmt.Println(timmy.name)
	// 这两行的输出结果完全是一模一样的，Go语言为结构体设置了自动取地址的特性
}

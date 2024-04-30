package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	timmy := &Person{
		name: "zhangweilong",
		age:  26,
	}
	fmt.Println(timmy.name)
	// 在解引用指针类型的结构体变量时，不用加上解引用操作符*
}

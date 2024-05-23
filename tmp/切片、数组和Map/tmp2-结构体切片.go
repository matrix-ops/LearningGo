package main

import "fmt"

type people struct {
	name    string
	age     int
	address string
}

func main() {
	var s []*people
	// 用var声明的切片长度和容量都为0
	fmt.Println(s)
	s = append(s, &people{
		name:    "zhangweilong",
		age:     25,
		address: "GuangDong",
	})
	fmt.Println(*s[0])
}

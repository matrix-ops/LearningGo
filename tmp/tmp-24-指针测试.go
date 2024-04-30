package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	name    string
	age     int
	address string
	people  struct {
		gender string
	}
}

func main() {
	var t *string
	s := "zhangweilong"
	t = &s
	fmt.Println(t)
	// 打印t里面存储的内存地址
	fmt.Println(&s)
	// 打印s的内存地址，也就是t存储的内存地址
	fmt.Println(t == &s)
	// true
	fmt.Println(reflect.TypeOf(t))
	// *string

	var e *Employee
	e = &Employee{
		name:    "zhangweilong",
		age:     10,
		address: "GuangDong",
		people: struct {
			gender string
		}{gender: "male"},
	}
	E := Employee{
		name:    "zhangweilong",
		age:     10,
		address: "GuangDong",
		people: struct {
			gender string
		}{gender: "male"},
	}
	fmt.Println(reflect.TypeOf(e))
	// *main.Employee
	fmt.Println(reflect.TypeOf(e.people))
	// struct { gender string }
	fmt.Println(reflect.TypeOf(E))

	S := new(string)
	// *string
	// 在内存中初始化一段内存地址，将这段内存地址里面存储的值，设置为string类型的零值
	// 将这个内存地址作为指针返回，因此S是一个*string类型的零值
	fmt.Println(S)
	//
	*S = "zhangweilong"
	fmt.Println(*S)
	// zhangweilong

}

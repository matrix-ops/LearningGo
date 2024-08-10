package main

import "fmt"

type Person struct {
	man
}
type man struct {
	name string
	age  int
	dick bool
}

func main() {
	p1 := Person{man{name: "zhangweilong", age: 25, dick: true}}
	p2 := p1
	fmt.Println(p1, p2)
	fmt.Println(p1.name)
	fmt.Println(p2.name)
	p2.name = "xiongmeiqi"
	// 更改新的结构体的同一个字段，不会影响老的结构体，下面输出的name字段的值将是xiongmeiqi
	fmt.Println(p1.name)
	fmt.Println(p2.name)
}

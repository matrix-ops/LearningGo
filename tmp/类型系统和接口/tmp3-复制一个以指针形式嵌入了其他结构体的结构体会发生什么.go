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
	fmt.Println(p1.name)
	fmt.Println(p2.name)
}

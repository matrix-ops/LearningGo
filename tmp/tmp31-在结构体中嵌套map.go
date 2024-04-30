package main

import "fmt"

type Visited struct {
	name    string
	visited map[string]int
}

func main() {
	v1 := Visited{name: "zhangweilong", visited: make(map[string]int)}
	// make函数的方式
	v1.visited["Google"] = 1000
	fmt.Println(v1)
	v2 := Visited{name: "zhangweilong", visited: map[string]int{"Baidu": 100}}
	// 字面量的方式
	fmt.Println(v2)
}

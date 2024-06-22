package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s)
	// []
	fmt.Println("s的长度是：", len(s), "s的容量是:", cap(s))
	// 0
	s1 := make([]int, 10)
	// 声明并初始化
	fmt.Println(s1)
	// [0 0 0 0 0 0 0 0 0 0]
	s2 := s
	fmt.Println(s2)
	// []
	var a [10]int
	fmt.Println(a)
}

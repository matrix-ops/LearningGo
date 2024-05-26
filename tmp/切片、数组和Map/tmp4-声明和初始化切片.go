package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s)
	// []
	s1 := make([]int, 10, 10)
	// 声明并初始化
	fmt.Println(s1)
	// [0 0 0 0 0 0 0 0 0 0]
	s2 := s
	fmt.Println(s2)
	// []
	var a [10]int
	fmt.Println(a)
}

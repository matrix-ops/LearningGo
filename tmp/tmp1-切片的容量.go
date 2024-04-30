package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3, 4, 5}
	fmt.Println(cap(a1), len(a1))
	// 5, 5
	a1 = append(a1, 6)
	fmt.Println(cap(a1), len(a1))
	// 10, 6
	// 当切片的长度超出初始容量，Go语言将把容量自动翻倍
}

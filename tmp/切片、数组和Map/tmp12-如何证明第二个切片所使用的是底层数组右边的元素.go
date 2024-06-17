package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[1:3]
	fmt.Println(len(newSlice), cap(newSlice))
	// 2, 4
	newSlice = append(newSlice, 50)
	fmt.Println(newSlice)
	// 2,3,50
	fmt.Println(slice)
	// 1,2,3,50,5
	// 说明两者共享同一个底层底层数组
	// 并且newSlice在slice的基础上创建时
	// 容量的剩余部分是底层数组右边没有被包括在创建newSlice时所使用的索引范围内的元素
}

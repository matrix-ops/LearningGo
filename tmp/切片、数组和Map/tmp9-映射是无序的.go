package main

import "fmt"

func main() {
	map1 := make(map[string]bool)
	map1["red"] = true
	map1["yellow"] = true
	map1["blue"] = false
	map1["green"] = false
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	fmt.Println(map1)
	fmt.Println(map1)
	fmt.Println(map1)
	// Go的map数据结构底层是一个hash表
	// 通过一个hash函数计算得出一个hash值，将hash值与桶的总数执行位运算（相对于取模来说性能更佳）来得到这个key应该被存储到哪个桶上
	// 桶是一个bmap结构体，结构体里面有三个关键字段：tophash、keys、values
	// 这三个字段都是数组，其中tophash用于存储hash结果的高八位，用于定位key存储于keys数组的哪个索引上
	// keys数组和values数组中某个元素的关系是一一对应的，例如如果一个key被确定存储于keys数组的索引2,那么它对应的值也存储于values数组的索引2
}

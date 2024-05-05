package main

// append只能用于给slice添加元素，不能给其他的数据结构添加元素

import (
	"fmt"
	"unsafe"
)

func main() {
	dwarfs1 := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter"}
	dwarfs2 := append(dwarfs1, "Saturn", "Uranus", "Neptune")
	fmt.Println(dwarfs2)
	fmt.Println(len(dwarfs1), cap(dwarfs1))
	fmt.Println(len(dwarfs2), cap(dwarfs2))
	var hex int = 0x123411E211227788
	result := hex & 255
	fmt.Println(result)
	fmt.Println(unsafe.Sizeof(hex))
	// 一个
	// 从数组生成切片
	var a6 = [5]int{1, 2, 3, 4, 5}
	s6 := a6[1:4]
	fmt.Println(s6)
	// 左闭右开区间
	// 前包后不包
	// 2,3,4
	fmt.Println(cap(s6), len(s6))
	// 4, 3
	// 容量之所以为4，是因为底层数组是2,3,4,5
	// 注意，如果是从一个数组中生成切片，并且切片的第一个元素不是数组的第一个元素
	// 那么切片的底层数组将是从切片的第一个元素在底层数组的位置，到底层数组的最后一个元素
	//

}

package main

import "fmt"

func main() {
	x := make([]int, 0) // 容量为 0
	x = append(x, []int{1, 2, 3, 4, 5}...)
	fmt.Println(cap(x)) // 观察容量的变化
	// 6
	// 之所以是6而不是5,是因为go编译器在执行底层数组重新分配容量的函数growslice中
	// 有一段内存对齐的代码
}

package main

import "fmt"

func main() {
	s1 := make([]int, 1, 5)
	s1 = []int{1, 2, 3, 4, 5}
	fmt.Println("s1的长度是：", len(s1), " s1的容量是", cap(s1))
	s1 = append(s1, 6)
	fmt.Println("s1的长度是：", len(s1), " s1的容量是", cap(s1))
	a1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := a1[0:5]
	// 从一个数组新建一个切片时，最好提前限制好新切片的容量
	// 提前限制好一个切片最大能有多大容量，会避免切片在后续扩容时，在原始数组上面分配到更大的容量，从而将原始数组原来的元素覆盖掉
	//
	fmt.Println("s2的长度是：", len(s2), " s2的容量是", cap(s2))
	s2 = append(s2, 60)
	fmt.Println("s2的长度是：", len(s2), " s2的容量是", cap(s2))
	s2 = append(s2, 60)
	fmt.Println(s2)
	fmt.Println(a1)
	fmt.Fprintln()
}

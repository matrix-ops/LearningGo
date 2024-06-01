package main

import "fmt"

func main() {
	s1 := []int{0, 1, 2, 3, 4, 5}
	s2 := s1[1:3]
	fmt.Println("在append之前，s2的容量是：", cap(s2))
	s2 = append(s2, 30)
	// s1[3]也会变成30
	// 两者共享同一个底层数组
	fmt.Println("在append之后，s2的容量是：", cap(s2))
	fmt.Println(s1)
}

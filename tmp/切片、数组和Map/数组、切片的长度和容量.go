package main

import "fmt"

func main() {
	var a1 = [2]string{"a"}
	fmt.Println("a1的长度是:", len(a1), "a1的容量是:", cap(a1))
	// 2 2
	// 数组的长度和容量都是不可变的
	// 即使在初始化的时候没有用光全部的容量，长度也等于容量
	var s1 = []string{"a", "b"}
	fmt.Println("s1的长度是:", len(s1), "s1的容量是:", cap(s1))
	s1 = append(s1, "c")
	fmt.Println(len(s1), cap(s1))
	// 3,4
	// a1 = append(a1, "b", "c")
	// 数组的长度和容量是不可变的
}

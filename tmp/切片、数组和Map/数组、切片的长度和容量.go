package main

import "fmt"

func main() {
	var a1 = [2]string{"a"}
	fmt.Println(len(a1), cap(a1))
	var s1 = []string{"a", "b"}
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, "c")
	fmt.Println(len(s1), cap(s1))
	// a1 = append(a1, "b", "c")
	// 数组的长度和容量是不可变的
}

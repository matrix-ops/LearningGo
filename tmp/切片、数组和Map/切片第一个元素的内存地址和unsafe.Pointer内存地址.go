package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s1 = []int{1, 2, 3}
	fmt.Println(&s1[0])
	ptr := unsafe.Pointer(&s1[0])
	fmt.Println(ptr)
	// 两者是一样的值

	// 切片本身是一个24字节（在64位系统中）的结构体，我们可以使用unsafe.Sizeof函数来验证这一点
	var s2 []int
	fmt.Println(unsafe.Sizeof(s2))
	// 指针8字节
	// 长度8字节
	// 容量8字节
}

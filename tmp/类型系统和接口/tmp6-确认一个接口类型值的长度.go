package main

import (
	"fmt"
	"unsafe"
)

func main() {
	interfaceSize := unsafe.Sizeof(interface{}(nil))
	fmt.Println(interfaceSize)
	// 在64位系统中，将打印16，表示长度是16个byte
	// 第一个字包含了一个指向内部表的指针，这个内部表称之为iTable
	// iTable存储了此接口所存储的值的类型信息(例如，Person类型）和值所对应的一组方法（在对应类型上定义的那一组能满足接口要求的方法）
	// 第二个字包含了一个指向值本身的指针
}

package main

import "fmt"

func main() {
	// 对空切片和nil切片调用内置函数例如append、len、cap，效果是一样的
	// 空切片和nil切片的区别在于空切片的地址指针不是空的，而是确实指向了一个底层数组
	// 而nil切片的地址指针是一个nil值
	// 声明空切片
	var s1 = []int{}
	fmt.Println(s1)
	s2 := make([]int, 0)
	s3 := []int{}
	fmt.Println(s2, s3)
	// 空切片的作用：当需要返回一个数据库查询结果，结果得到0条满足条件的记录时，可以使用空切片

	// 声明nil切片
	var s4 []int
	// 只声明不初始化得到的就是一个nil切片
	fmt.Println(s4)
	// nil切片的实际意义
	// 当一个函数的返回值是一个切片，但是在函数内部发生了错误导致此切片的值不正常的时候，可以通过返回一个nil切片来通知调用者函数内部调用过程发生了错误
}

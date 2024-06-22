package main

import "fmt"

func main() {
	// var array1 [2][2]int
	// // 声明了变量就会被自动初始化为对应类型的零值，指针的零值是nil
	// // 之所以一个指针变量只声明不显式初始化然后直接使用的话会报错
	// // 是因为发生了空指针解引用
	// // 解引用nil指针会发生报错
	// array1[1] = [2]int{1, 2}
	// fmt.Println(array1)
	var array1 [2]int
	fmt.Println(array1)
	// [0 0]
}

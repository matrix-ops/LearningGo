package main

import "fmt"

func main() {
	// var array1 [2][2]int
	// // 声明了变量就会被自动初始化为对应类型的0值
	// // 之所以一个指针变量声明但是不显式初始化，直接使用的话会报错
	// // 是因为发生了空指针解引用
	// // 指针变量的被声明出来之后它也被初始化为对应的零值，零值就是一个空指针nil
	// // 解引用nil指针会发生报错
	// array1[1] = [2]int{1, 2}
	// fmt.Println(array1)
	var array1 [2]int
	fmt.Println(array1)
	// [0 0]
}

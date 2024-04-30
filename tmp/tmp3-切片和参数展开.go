package main

import (
	"fmt"
	"reflect"
)

func PrintlnNumbers(numbers ...int) []int {
	fmt.Println(reflect.TypeOf(numbers))
	// 打印传参的类型，在这个例子中，PrintNumbers函数的传参会被存储为一个int切片
	newNumbers := make([]int, len(numbers))
	// 初始化一个int切片，长度为传参的个数
	for i := range numbers {
		newNumbers[i] = numbers[i] + 100
		// 遍历传参
		// 注意，通常情况下，用range函数遍历切片会返回两个值，第一个值是索引
		// 第二个值是元素的值
		// 如果只写一个，那就只返回索引

	}
	return newNumbers
}
func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(PrintlnNumbers(slice1...))
}

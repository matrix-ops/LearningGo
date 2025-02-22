package main

import "fmt"

func main() {
	// 外部函数返回一个闭包
	add := func(x int) func(int) int {
		return func(y int) int {
			return x + y
		}
	}

	// 创建一个闭包，x 绑定为 10
	add10 := add(10)
	fmt.Println(add10(0))

	// 调用闭包时，y 为 5
	fmt.Println(add10(5)) // 输出 15
}

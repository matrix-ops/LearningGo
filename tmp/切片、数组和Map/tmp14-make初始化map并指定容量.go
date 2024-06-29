package main

import "fmt"

func main() {
	m1 := make(map[string]string, 15)
	fmt.Println(len(m1))
	// 0
	// 注意,使用make函数创建map的时候第二个参数只能是容量，指定的并非是长度
}

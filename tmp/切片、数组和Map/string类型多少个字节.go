package main

import "fmt"

func main() {
	var s1 = "zhangweilong"
	// 12
	fmt.Println(len(s1))
	var s2 = "123"
	// 3
	fmt.Println(len(s2))
	// 一个字母一个字节
	// 一个汉字三个字节
	var s3 = "张威龙"
	fmt.Println(len(s3))
	// 9
}

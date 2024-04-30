package main

import "fmt"

func main() {
	bolden := "Charles Bolden"
	charles := bolden
	// fmt.Println(charles == bolden)
	//	go语言不同类型不能对比
	fmt.Println(charles == bolden)
	// true
	fmt.Println(&charles == &bolden)
	// false
}

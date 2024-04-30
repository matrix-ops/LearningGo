package main

import (
	"fmt"
	"reflect"
)

func printAddress(a *int) {
	fmt.Println("参数a的值是一个内存地址，内存地址里面的值是：", *a)
	// 这里说的是内存地址里面的值
	// 如果想获取内存地址，直接打印a，不要使用*a即可
	fmt.Println("参数a的值是一个内存地址，是：", a)
}
func main() {
	answer := 43
	fmt.Println("answer的内存地址是：", &answer)
	// &取地址操作符通常作用于普通变量名之前，用于获取变量的内存地址
	address := &answer
	// 将answer的内存地址赋值给一个变量
	// 请记住，指针也可以是一种变量，当变量的值是内存地址是，这个变量就可以称之为指针
	// 上面的情况下，address将被自动初始化为一个int类型的指针
	// 值为answer的内存地址
	fmt.Println(reflect.TypeOf(address))
	printAddress(address)
}

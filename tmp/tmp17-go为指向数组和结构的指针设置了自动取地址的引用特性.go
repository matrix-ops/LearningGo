package main

import (
	"fmt"
	"reflect"
)

func main() {
	power := &[2]int{1, 2}
	power[0] += 10
	fmt.Println(power[0])
	// 11
	fmt.Println(reflect.TypeOf(power))
	// *[2]int
	// 不仅如此，所有接收者为指针类型的方法，Go语言也提供了自动取地址的特性
	// 把一个值变量传递给一个只接收指针类型的方法，Go语言会自动在传递的时候引用值变量的地址
	// 再进行传递
}

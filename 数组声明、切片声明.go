package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a [10]string
	var b = [...]int{1, 2, 3, 4, 5, 6}
	c := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	var s string
	var sep string = " "
	a[0] = "a"
	fmt.Println(len(a))
	fmt.Println(a[0])
	fmt.Println(b)
	for _, args := range c {
		s += args + sep
	}
	fmt.Println("s的值是", s)
	// for i, args := range b {
	// 	fmt.Printf("%d\t%d\n", i, args)
	// }
	slice1 := []int{10, 20, 30, 40, 50, 60}
	for i, element := range slice1 {
		fmt.Printf("%d %d\n", i, element)
	}
	fmt.Println(reflect.TypeOf(slice1))
	// slice1是一个切片
}

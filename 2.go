package main

import (
	"fmt"
	"reflect"
)

func main() {
	str1 := "Will Love You"
	s1 := []string{"test", "output"}
	a1 := [2]string{"test", "output"}
	m1 := map[string]string{
		"name": "zhangweilong",
		"age":  "18",
	}
	// 初始化的时候cap是2
	s1 = append(s1, "Ok")
	fmt.Println(cap(s1), len(s1))
	fmt.Println(reflect.TypeOf(s1))
	fmt.Println(reflect.TypeOf(a1), len(a1))
	fmt.Println(reflect.TypeOf(m1), reflect.TypeOf(str1))
}

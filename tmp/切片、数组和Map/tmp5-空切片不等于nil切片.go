package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a1 [3]int
	var a2 [3]int
	fmt.Println(a1 == a2)
	var s1 []int
	var s2 []int
	s3 := []int{}
	fmt.Println(reflect.DeepEqual(s1, s2))
	// true
	fmt.Println(reflect.DeepEqual(s1, s3))
	// false
}

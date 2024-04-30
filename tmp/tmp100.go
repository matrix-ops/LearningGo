package main

import "fmt"

type testStruct struct {
	name string
	age  int
	addr string
}

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	// fmt.Print(testStruct{})
	for i, err := range s1 {
		fmt.Println(i, err)
	}
}

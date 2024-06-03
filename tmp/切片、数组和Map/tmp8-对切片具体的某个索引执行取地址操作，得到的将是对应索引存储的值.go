package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	fmt.Println(&s1[0])
	p1 := &s1[0]
	*p1 = 10
	fmt.Println(s1)
	// [10,2,3]
}

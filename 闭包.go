package main

import "fmt"

func incr(x int) func() int {
	return func() int {
		x++
		return x
	}
}
func main() {
	s := incr(2)
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(incr(2)())
}

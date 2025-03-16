package main

import "fmt"

func main() {
	add := func(x int) func(y int) int {
		return func(y int) int {
			return x + y
		}
	}
	add10 := add(10)
	fmt.Println(add10(10))
}

package main

import (
	"fmt"
	"sort"
)

func adder(a, b int) int {
	return a + b
}
func main() {
	var fn func(a, b int) int
	fn = adder
	fmt.Println(fn(1, 2))
	sort.Slice()
}

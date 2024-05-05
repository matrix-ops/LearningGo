package main

import "fmt"

type typeAdder func(x int) int

func adder(y int) typeAdder {
	return func(x int) int {
		return x + y
	}
}

func main() {
	adderTest := adder(2)
	fmt.Println(adderTest(3))
}

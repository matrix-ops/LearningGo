package main

import "fmt"

func main() {
	fn := func() {
		fmt.Println("Fuck you")
	}
	fn()
	fmt.Printf("%T\n", fn)
}

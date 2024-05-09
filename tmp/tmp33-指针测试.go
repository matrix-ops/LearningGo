package main

import "fmt"

func main() {
	var p *int
	i := 10
	p = &i
	fmt.Println(*p)
	fmt.Println(p)
	i = 12
	fmt.Println(*p)
	fmt.Println(p)
	*p = 42
	fmt.Printf("%T\n", p)
	answer := 42
	address := &answer
	fmt.Printf("%T\n", address)
}

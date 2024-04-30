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
}

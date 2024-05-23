package main

import "fmt"

func main() {
	fmt.Println(new(int))
	a := new(int)
	*a = 1
	fmt.Println(*a)

}

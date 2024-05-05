package main

import "fmt"

func main() {
	var number int = 10
	switch number > 9 {
	case true:
		fmt.Println("True")
	case false:
		fmt.Println("False")
	}
}

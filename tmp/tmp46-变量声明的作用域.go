package main

import "fmt"

func main() {
	runner := 2
	if true {
		runner := 1
		fmt.Println(runner)
	}
	fmt.Println(runner)
}

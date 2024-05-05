package main

import "fmt"

func main() {
	f1 := fmt.Sprintln("I am your father, my son")
	fmt.Println(f1)
	f2 := fmt.Sprintf("I love you %s", f1)
	fmt.Println(f2)
}

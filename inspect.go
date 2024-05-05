package main

import "fmt"

func main() {
	var year int64 = 2018
	var1 := "love you mother"
	var2 := 12
	var3 := 0.345
	var4 := true
	fmt.Printf("type: %T, %[1]v\n", year)
	fmt.Printf("type: %T, %[1]v\n", var1)
	fmt.Printf("type: %T, %[1]v\n", var2)
	fmt.Printf("type: %T, %[1]v\n", var3)
	fmt.Printf("type: %T, %[1]v\n", var4)
}

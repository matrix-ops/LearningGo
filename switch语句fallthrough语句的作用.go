package main

import "fmt"

func main() {
	k := 2
	switch k {
	case 1:
		fmt.Println("This is case 1")
		fallthrough
	case 2:
		fmt.Println("This is case 2")
		fallthrough
	case 3:
		fmt.Println("This is case 3, but might not get here without fallthrough.")
	default:
		fmt.Println("This is the default case")
	}
}

// 会输出
// This is case 2
// This is case 3, but might not get here without fallthrough.

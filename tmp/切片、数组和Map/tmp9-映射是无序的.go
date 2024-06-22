package main

import "fmt"

func main() {
	map1 := make(map[string]bool)
	map1["red"] = true
	map1["yellow"] = true
	map1["blue"] = false
	map1["green"] = false
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	fmt.Println(map1)
	fmt.Println(map1)
	fmt.Println(map1)
}

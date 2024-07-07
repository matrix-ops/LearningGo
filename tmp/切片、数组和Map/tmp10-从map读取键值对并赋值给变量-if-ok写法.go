package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["age"] = 19
	msg, ok := m1["age"]
	if ok {
		fmt.Println(ok)
	}
	// if ok写法，如果在m1里面存在"age"这个键，ok将被设置为true
	fmt.Println(msg)
}

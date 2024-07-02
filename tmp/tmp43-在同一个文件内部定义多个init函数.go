package main

import "fmt"

func init() {
	fmt.Println("这是第一个init函数")
}
func init() {
	fmt.Println("这是第二个init函数")
}
func main() {
	fmt.Println("这是main函数")
}

func speaker(s string) {
	fmt.Println(s)
}

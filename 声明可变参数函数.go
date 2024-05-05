package main

import "fmt"

func Terraform(prefix string, words ...string) {
	for i := range words {
		words[i] = prefix + " " + words[i]
	}
}
func main() {
	planets := []string{"Earth", "Mars"}
	Terraform("new", planets...)
	fmt.Println(planets)
}

// 传递给函数的可变长形参如果是切片展开后进行传入的，那么针对每个参数的修改将直接作用于切片，作用类似于指针传递

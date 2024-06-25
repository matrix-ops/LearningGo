package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	prints := []func(){}
	for _, v := range arr {
		fmt.Println(&v)
		prints = append(prints, func() { fmt.Println(v) })
		// 在go 1.22之前，这里应该打印的是同一个内存地址
		// 因为会创建一个临时变量v，之后的每次迭代都只是将切片的值赋值给了这个临时变量
		// go 1.22之后，每次循环都会重新创建一个变量v，因此每次迭代里面v的内存地址都是不一样的
	}
	for _, p := range prints {
		p()
	}
}

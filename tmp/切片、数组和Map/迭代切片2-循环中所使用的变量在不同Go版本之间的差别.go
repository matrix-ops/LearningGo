package main

import "fmt"

func main() {
	var prints []func()
	for i := 0; i < 5; i++ {
		prints = append(prints, func() { fmt.Println(i); fmt.Println(&i) })
		// 在Go 1.22之前，此for循环内部的i是同一个变量
		// Go1.22之后，每次循环都会重新创建一个变量i,第一次循环会初始化一个int类型的i,之后的post语句（i++）都会在
		// 前一次迭代的i值的基础上执行，例如，在这个例子中，第二次迭代的post语句是将i设置成了3（i=1;i++;i++）
		// 此外，在这个循环的例子中prints切片里面所存储的函数值，其实都是闭包函数，里面留存了创建闭包时i的内存地址和值
		fmt.Println(&i)
		i++
	}
	for _, print := range prints {
		print()
	}
}

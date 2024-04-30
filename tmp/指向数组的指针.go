package main

import "fmt"

type grid [9][9]int

func main() {
	var p *[8]int
	p = &[8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(*p)
	var g *grid
	g[1][2] = 1
	// 会空指针引用，因为这个元素并不存在，你也只是声明了这个数组但是并没有初始化
	// 注意，数组不能用make函数来声明
	fmt.Println(g)
}

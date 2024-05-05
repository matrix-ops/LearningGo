package main

import "fmt"

type Person37 struct {
	age  int
	name string
}

func (p *Person37) birthday() {
	p.age++
	// 这里不用解引用是因为Go语言会自动对结构体指针做解引用
}

type address string

func (a *address) printAddress() {
	fmt.Println(a)
	// 字符串类型并不能自动解引用
	fmt.Println(*a)
	// 正常解引用
}
func main() {
	p1 := Person37{
		age:  25,
		name: "zhangweilong",
	}
	p1.birthday()
	fmt.Println(p1)

	a1 := address("GuangDong")
	a1.printAddress()
}

// 结构体指针作为接收者的时候，值也可以直接调用birthday方法
// 字符串指针作为接收者的时候，字符串值也可以直接调用printAddress方法

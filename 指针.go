package main

import "fmt"

type testInterface interface {
	address() string
}

// 模拟error接口

type testString struct {
	name string
}

// 模拟errorString

func (ss *testString) address() string {
	return ss.name
}

// 模拟Error()方法

func printAddress(text string) testInterface {
	return &testString{text}
}

// 返回一个结构体字面量，内容是text
// testString类型实现了testInterface接口
// 模拟New函数

func main() {
	Address := printAddress("GuangDong Sheng")
	address := printAddress("GuangDong Sheng")
	fmt.Printf("%v\n", Address.address())
	fmt.Printf("%v\n", address.address())

	var i int = 1
	var p *int = &i
	// p是一个int类型的指针，它只能接受指向int类型的值的指针
	fmt.Println(*p)
}

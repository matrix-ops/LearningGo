package main

import "fmt"

type number struct {
	value int
	valid bool
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

func (n number) String() string {
	// 类型的String方法用来定制输出
	if !n.valid {
		return "not set"
	}
	return fmt.Sprintf("%d", n.value)
}

func main() {
	n := newNumber(32)
	fmt.Println(n)
	n1 := number{}
	fmt.Println(n1)
}

package main

import (
	"fmt"
	"strings"
)

type fucker interface {
	talkName() string
}

func (p Person) talkName() string {
	return strings.ToUpper(p.name)
}

type Person struct {
	age  int
	name string
}

func NewPerson(name string, age int) *Person {
	p := &Person{
		name: name,
		age:  age,
	}
	return p
}

func (p *Person) GrowUp(i int) {
	p.age = i
}

func main() {
	p := Person{age: 1, name: "zhangweilong"}
	fmt.Println(p)
	p.GrowUp(18)
	fmt.Println(p)
	var t1 fucker
	t1 = p
	fmt.Println(t1.talkName())
	// p是一个值，因此它满足了fucker接口
	p1 := &Person{age: 1, name: "dolphin"}
	t1 = p1
	fmt.Println(t1.talkName())
	// p1是一个指针，它同样满足了fucker接口
}

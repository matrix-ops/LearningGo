package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return d.name
}

type Cat struct {
	name string
}

func (c Cat) Speak() string {
	return c.name
}

func main() {
	var s1 Speaker
	var d Dog
	var c Cat
	d.name = "Dog"
	c.name = "Cat"
	s1 = d
	fmt.Println(s1.Speak())
	// Dog
	s1 = c
	fmt.Println(s1.Speak())
	// Cat
	// 这就是多态
}

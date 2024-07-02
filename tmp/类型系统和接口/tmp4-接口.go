package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (d Dog) Speak() string {
	return d.name
}

func (c Cat) Speak() string {
	return c.name
}
func Speak(s Speaker) string {
	return s.Speak()
	// 注意，接口是不允许有自己的方法的，也就是说不能声明一个接收者为接口的方法
}

func main() {
	d1 := new(Dog)
	d1.name = "Dog"
	fmt.Println(Speak(d1))
	fmt.Println(d1.Speak())
}

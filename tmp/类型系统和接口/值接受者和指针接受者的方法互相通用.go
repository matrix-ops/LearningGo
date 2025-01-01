package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) SayHello() {
	fmt.Println("Hello World, I am " + p.name)
}

func (p *Person) ChangeAge(newAge int) {
	p.age = newAge
}

func main() {
	var p1 Person = Person{
		name: "zhangweilong",
		age:  7,
	}
	var p2 *Person = &Person{
		name: "xiongmeiqi",
		age:  100,
	}
	p2.SayHello()
	// 指针调用值接受者方法
	p1.ChangeAge(18)
	// 值调用指针接受者方法
	fmt.Println(p1)

}

package main

import "fmt"

type person struct {
	superpower string
	age        int
}

func birthdays(p *person) {
	p.age++
	// 这里不用写成（*p).age++，因为go语言会自动解引用
}

func main() {
	p := &person{
		superpower: "eat",
		age:        19,
	}
	birthdays(p)
	fmt.Printf("%+v", p)
}

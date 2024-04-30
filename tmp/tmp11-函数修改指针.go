package main

import "fmt"

type person struct {
	superpower string
	age        int
}

func birthdays(p *person) {
	p.age++
}

func main() {
	p := &person{
		superpower: "eat",
		age:        19,
	}
	birthdays(p)
	fmt.Printf("%+v", p)
}

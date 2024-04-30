package main

import "fmt"

type person struct {
	superpower string
	age        int
}

func (p *person) birthdays() {
	p.age++
}
func main() {
	p := &person{
		superpower: "sex",
		age:        19,
	}
	p.birthdays()
	fmt.Printf("%+v", p)
}

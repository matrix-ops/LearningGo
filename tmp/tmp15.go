package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) birthdays() *person {
	p.age++
	return nil
}

func main() {
	var p *person
	fmt.Println(p)
	p.birthdays()
}

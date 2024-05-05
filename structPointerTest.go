package main

import "fmt"

type person struct {
	name    string
	age     int
	address string
}

func (p *person) Say(words string) {
	p.name = p.name + words
}

func main() {
	p := person{
		name:    "zhangweilong",
		age:     18,
		address: "GuangDong GuangZhou",
	}
	p.Say(" Fuck you")
	fmt.Println(p)
	a := 60
	b := 0x15
	c := 0x51
	fmt.Println(b % a)
	fmt.Println(c % a)

}

package main

import "fmt"

type Person38 struct {
	age     int
	name    string
	Address string
}

func newPerson(age int, name string) *Person38 {
	p := Person38{age: age, name: name}
	return &p
}

func (p *Person38) ageAdd() {
	p.age += 1
}

func main() {
	p1 := newPerson(18, "zhangweilong")
	fmt.Println(p1)
	p2 := Person38{
		age:  18,
		name: "zhangweilong",
	}
	p2.ageAdd()
	fmt.Println(p2)
}

package main

import "fmt"

type Fucker interface {
	fuck()
	suck()
}

type xiongmeiqi struct {
	name string
	age  int
}

func (x xiongmeiqi) fuck() {
	fmt.Println(x.name, x.age, "fucking")
}
func (x xiongmeiqi) suck() {
	fmt.Println(x.name, x.age, "sucking")
}
func main() {
	x := xiongmeiqi{
		name: "xiongmeiqi",
		age:  28,
	}
	var f Fucker
	f = x
	f.suck()
	f.fuck()
}

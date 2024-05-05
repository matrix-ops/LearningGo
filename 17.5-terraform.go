package main

import "fmt"

type Planets []string

func (p Planets) terraform() Planets {
	for i := range p {
		p[i] = "New_" + p[i]
	}
	return p
}

func main() {
	var p = Planets{"Mercury", "Venus", "Earth"}
	// []Planets{} 是错误写法
	// 这样写的意思是一个由Planets{}组成的切片
	p = p.terraform()
	// p = Planets.terraform(p) 是同样写法
	fmt.Println(p)
}

package main

import "fmt"

func Terraform(p *[8]string) {
	// 这是类型声明，表明这是一个指针
	for i := 0; i < len(p); i++ {
		p[i] = "New " + p[i]
	}
	// 这是解引用
}
func main() {
	plants := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	Terraform(&plants)
	fmt.Println(plants)
}

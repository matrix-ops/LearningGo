package main

import "fmt"

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	terrestrial := planets[0:4]
	// worlds := append(terrestrial, "Cers")
	fmt.Println(append(terrestrial, "Cers"))
	// 切片是引用类型，因此即使此处的append没有将返回值复制给任何新的变量，底层的数组仍然被修改了
	fmt.Println(planets)
}

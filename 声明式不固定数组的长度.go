package main

import "fmt"

func terraform(planets [8]string) [8]string {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
	return planets
}
func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	// 不固定长度的数组，这仍然表明这是一个数组，而非一个切片，中括号里面有三个点的就是可变长数组
	planets = terraform(planets)
	fmt.Println(planets)
}

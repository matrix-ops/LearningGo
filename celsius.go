package main

import "fmt"

func main() {
	type celius float64
	var temperature celius = 20.0000
	fmt.Println(temperature)
	fmt.Printf("%.2f", temperature)
	// 20.00
}

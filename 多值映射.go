package main

import "fmt"

func main() {
	var testMap = map[string][3]float32{
		"Name": {123.0, 345.0, 456.0},
	}
	fmt.Println(testMap)
}

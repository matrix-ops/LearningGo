package main

import "fmt"

func main() {
	//var soup map[string]int
	//soup = map[string]int{"onicon": 1}
	//test_soup, ok := soup["oncon"]
	//if ok {
	//	fmt.Println(soup["onicon"])
	//}
	//fmt.Println("test_soup", test_soup)
	var soup map[string]int
	fmt.Println(soup == nil)
	soup = map[string]int{"onion": 1}
	measurement, ok := soup["onion"]
	if ok {
		fmt.Println(measurement)
	}
	for ingredient, measurement := range soup {
		fmt.Println(ingredient, measurement)
	}
}

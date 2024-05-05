package main

import "fmt"

func main() {
	var number int8
	arr := []int{2, 3, 5, 7, 11, 13}
	slic := append(arr, 17)
	fmt.Println(slic, cap(slic))
	fmt.Println(arr)
	fmt.Println(number)
}

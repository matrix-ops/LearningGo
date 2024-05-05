package main

import "fmt"

func main() {
	dwarfSlice := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter"}
	dwarfArray := [3]string{"Saturn", "Uranus", "Neptune"}
	fmt.Printf("%T\n", dwarfSlice)
	fmt.Printf("%T\n", dwarfArray)
	dwarfSlice = append(dwarfSlice, "Sun")
	fmt.Println(len(dwarfSlice), cap(dwarfSlice))
	// 切片的长度超过初始数组的长度时，它的底层数组会被替换成两倍于初始数组长度的新数组
}

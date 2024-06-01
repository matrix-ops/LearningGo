package main

import "fmt"

func main() {
	var array1 [3]*string
	array2 := [3]*string{new(string), new(string), new(string)}
	fmt.Println(array1)
	*array2[0] = "Yellow"
	*array2[1] = "Blue"
	*array2[2] = "Red"
	array1 = array2
	// 此时array2内部存储的实际上是内存地址
	// 赋值会把三个元素的内存地址都赋值给array1
	// 修改array1里面的内存地址所存储的值，结果会直接在array2上也能看到
	fmt.Println(array1, array2)
	// 在这里，两个数组中所存储的内存地址是完全一致的
	array1[0] = new(string)
	// 给array1的第一个元素分配一个新的内存地址
	// 此时修改array1的第一个元素里面的值，不会直接表现在array2的第一个元素上
	// 因为他们已经不是同一个内存地址
	*array1[0] = "White"
	fmt.Println(array1, array2)
	// 可以看到两个元素的内存地址已经不一样了
	for v, _ := range array1 {
		fmt.Println(*array1[v], *array2[v])
	}
}

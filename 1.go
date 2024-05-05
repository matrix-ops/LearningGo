package main

import "fmt"

var j int = 1

func showMemoryAddress(x *int) {
	// 指针的值必然是一个内存地址
	// 取地址操作符&可以取到一个变量的内存地址，那么这样在变量前面加取地址操作符，就可以当成指针的值来用
	//
	// 注意这里的x表示在这个函数内部它是一个指针，它的值是函数调用者传递进来的
	// 因此在函数内部的x有自己的内存地址

	// 指针是一种变量，它的值是一个内存地址，至于是谁的内存地址，那是另外一回事
	// 同时指针自己作为一个变量，也存在自己的内存地址
	fmt.Println("在showMemoryAddress函数内部，x的内存地址是:", &x)
	fmt.Println("你给showMemoryAddress函数传入参数的内存地址是:", x)
	return
}
func main() {
	var i int = 1
	showMemoryAddress(&i)
	fmt.Println("变量i的内存地址是：", &i)

	map1 := map[string]int{
		"zhangweilong": 10,
		"xiongmeiqi":   11,
	}
	// 初始化一个map变量，它的键是string，值是int
	fmt.Println(map1)
	array1 := [2]string{"zhangweilong", "xiongmeiqi"}
	// 初始化一个数组,长度为2，容量为2，元素类型为string
	fmt.Println(array1)
	slice1 := []string{"zhangweilong", "xiongmeiqi"}
	// 初始化一个切片，长度为2，容量为2，元素类型为string
	fmt.Println("slice1的值是: ", slice1, "slice1的长度是：", len(slice1), "slice1的容量是: ", cap(slice1))
	slice1 = append(slice1, "zhangyujing")
	fmt.Println("新增一个元素之后，slice1的长度和容量为：", len(slice1), cap(slice1))
	slice1 = append(slice1, "huangyuting")
	fmt.Println("再新增一个元素之后，slice1的长度和容量为：", len(slice1), cap(slice1))
	slice1 = append(slice1, "huangshuting")
	fmt.Println("由于超出了容量为4的限制，现在slice1的长度和容量为：", len(slice1), cap(slice1))

	var s = "Hello Mother Fucker"
	var p1 *string
	p1 = &s
	// p1是一个指针，他能接受一个值为string的变量的内存地址作为它自己的值
	fmt.Println("p1内部存储的内存地址为:", p1)
	fmt.Println("p1自己的内存地址为:", &p1)
	fmt.Println("p1存储的内存地址解析之后为:", *p1)
}

package main

import "fmt"

func main() {
	m1 := make(map[string]int, 1)
	m1["zhangweilong"] = 1
	fmt.Println("一个元素的时候：", len(m1))
	m1["huangyuting"] = 2
	fmt.Println("两个元素的时候", len(m1))
	m1["huangshuting"] = 3
	fmt.Println("三个元素的时候", len(m1))
	fmt.Printf("%+v", m1)
	// 结论是如果map初始化的时候给出的容量不够的话，后续新增的元素会自动扩容map
	error
}

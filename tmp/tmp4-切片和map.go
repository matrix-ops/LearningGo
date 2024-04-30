package main

import "fmt"

type EmployeeInfo struct {
	name    string
	address string
	age     int
}

func main() {
	intSlice1 := make([]int, 10, 10)
	// 初始化一个int切片，长度和容量都是10
	fmt.Println(intSlice1)
	stringSlice1 := make([]string, 10, 10)
	// 初始化一个string切片，长度和容量都是10
	fmt.Println(stringSlice1)
	floatSlice1 := make([]float64, 10, 10)
	// 初始化一个float64切片，长度和容量都是10
	fmt.Println(floatSlice1)
	JiDiEmployeeInfo := make(map[string]EmployeeInfo)
	// 初始化一个map结构，键为string类型，值为EmployeeInfo类型，也就是一个结构体
	// 注意，append函数不能操作map数据结构，如果想往map里面新增元素，只能使用如下方式：
	JiDiEmployeeInfo["Operator"] = EmployeeInfo{
		name:    "zhangweilong",
		address: "GuangDong GuangZhou",
		age:     27,
	}
	fmt.Println(JiDiEmployeeInfo)
}

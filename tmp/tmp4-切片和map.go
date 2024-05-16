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
	intSlice2 := make([]int, 10)
	// 如果不指明容量，默认就是10，和长度保持一致，如果初始容量不需要大于长度
	// 就没有必要指明第三个参数
	fmt.Println(intSlice1)
	// 0 0 0 0 ...
	// 默认会使用类型的零值来填充
	fmt.Println(cap(intSlice2))
	stringSlice1 := make([]string, 10, 10)
	// 初始化一个string切片，长度和容量都是10
	fmt.Println(stringSlice1)
	// 空字符填充的切片
	floatSlice1 := make([]float64, 10, 10)
	// 初始化一个float64切片，长度和容量都是10
	fmt.Println(floatSlice1)
	JiDiEmployeeInfo := make(map[string]EmployeeInfo)
	// 初始化一个map结构，键为string类型，值为EmployeeInfo结构体类型
	// 注意，append函数不能操作map数据结构，如果想往map里面新增元素，只能使用如下方式：
	JiDiEmployeeInfo["Operator"] = EmployeeInfo{
		name:    "zhangweilong",
		address: "GuangDong GuangZhou",
		age:     27,
	}
	fmt.Println(JiDiEmployeeInfo)
}

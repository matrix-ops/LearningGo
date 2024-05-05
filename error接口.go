package main

import "fmt"

type testError1 string
type testError2 string

func (Err testError1) Error() string {
	return string(Err)
}
func (Err testError2) Error() string {
	return string(Err + "哈哈哈")
}

func main() {
	var s error
	// error是一个接口类型，此时s是一个接口
	var t testError1
	// testError1是一个自定义类型，基础类型是string，它实现了Error方法，因此可以被赋值给error接口类型
	t = "我成功理解了什么是接口"
	s = t
	// 接口类型的变量被赋值了一个具体的值
	fmt.Println(s.Error())
	// 此时打印的就是"我成功理解了什么是接口"
	s = testError2("我成功理解了什么是接口")
	fmt.Println(s)
	// 打印时默认就会调用error接口类型的Error()方法
	// 此时打印的就是"我成功理解了什么是接口哈哈哈"
	// 这几行代码充分展现了接口类型的灵活性，任意实现了接口定义的类型的值都可以随时赋值给接口变量

//
// type testError1 string
// type testError2 string
//
// func (Err testError1) Error(str string) string {
// 	return str
// }
// func (Err testError2) Error(str string) string {
// 	return str + "哈哈哈哈"
// }
//
// func main() {
// 	var s error
// 	var t1 testError1
// 	var t2 testError2
// 	s = t1
// 	// 这样之所以不行是因为testError1类型的Error方法签名和error接口中定义的不一样
// 	fmt.Println(s.Error("我成功理解了什么是接口"))
// 	s = t2
// 	fmt.Println(s.Error("我成功理解了什么是接口"))
// }

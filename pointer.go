package main

import "fmt"

type testStruct struct {
	s string
}

var (
	err = blueprint("Test")
)

func (t *testStruct) Error() string {
	return t.s + " Ha Ha"
}

func blueprint(x string) error {
	return &testStruct{x}
}

func main() {
	superpowers := &[3]string{"flight", "invisibility", "super strength"}
	fmt.Println("superpowers的第一个元素是：", (*superpowers)[0])
	var home *string
	canada := "Canada"
	home = &canada
	fmt.Printf("canada变量的内存地址 %v\n", &canada)
	fmt.Printf("home里面存储的内存地址：%v\n", home)
	// 如果不带*，就没办法解码，将直接打印指针内部所存储的那个内存地址
	fmt.Printf("home里面存储的内存地址，解码出来之后是什么: %v\n", *home)
	// *号用于解码指针变量内部所存储的内存地址值
	fmt.Printf("home自己的内存地址: %v\n", &home)
	// 同时，指针自己也有一个内存地址值
	// 	var test1 testStruct
	// 	test1.s = "Test Struct"
	// 	fmt.Println(test1.Error())

	var x *testStruct
	var testx testStruct = testStruct{
		"zhangweilong",
	}
	x = &testx
	fmt.Println(x)
	// 在上面，给testStruct这个类型定义了一个Error接口，fmt.Println会分别调用Error和Stringer方法，而且Error方法的优先级会比
	// Stringer更高，因此你会看到这里并没有显式的调用Error方法，结果Error方法还是被调用了
	fmt.Println(x.s)
}

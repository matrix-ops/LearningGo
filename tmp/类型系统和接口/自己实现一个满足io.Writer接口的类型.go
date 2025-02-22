package main

import (
	"fmt"
	"io"
)

type WriterTest struct {
	Name string
}

func (w *WriterTest) Write(p []byte) (n int, err error) {
	w.Name += string(p)
	return len(p), nil
}

func main() {
	var w WriterTest
	w.Name = "zhangweilong"
	var writer io.Writer = &w
	// 只要实现了Write(p []byte) (n int, err error)这个方法
	// 就实现了io.Writer接口
	// 此时i这个接口变量里面存储是w的内存地址
	fmt.Println("i's value is ", writer)
	// 打印io.Writer类型变量i的值
	fmt.Fprintf(writer, "haha")
	// Fprintf的第一个参数必须是io.Writer
	// Fprintf会调用writer的Write方法，因此"haha"会被写入到w.name
	fmt.Println(writer)
	// zhangweilonghaha

	// 不能直接通过io.Writer去访问实际类型的值
	// 需要使用类型断言
	// 类型断言
	if wt, ok := writer.(*WriterTest); ok {
		fmt.Println((*wt).Name)
	} else {
		fmt.Println("类型断言失败")
	}

}

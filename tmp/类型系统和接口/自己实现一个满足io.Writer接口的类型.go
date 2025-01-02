package main

import (
	"fmt"
	"io"
)

type WriterTest struct {
	name string
}

func (w *WriterTest) Write(p []byte) (n int, err error) {
	w.name += string(p)
	return len(p), nil
}

func main() {
	var w WriterTest
	w.name = "zhangweilong"
	var i io.Writer = &w
	// 此时i这个接口变量里面存储是w的内存地址
	fmt.Println("i's value is ", i)
	fmt.Fprintf(i, "haha")
	fmt.Println(w.name)
}

package main

import "fmt"

type ErrorTest struct {
	name string
	xing string
}

func (e *ErrorTest) Error() string {
	return e.xing
}

// 当ErrorTest类型的变量被当做error接口使用时，Error()方法无需显示调用，将自动被执行

func ErrorNew(text string) error {
	return &ErrorTest{name: text, xing: text + "haha"}
}
func main() {
	e := ErrorNew("zhang")
	fmt.Println(e)
}

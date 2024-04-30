package main

import (
	"fmt"
	"reflect"
	"strings"
)

type talKer interface {
	talking() string
}

type commander interface {
	call() string
	commanding(i string) string
}
type stripper string

func (s stripper) call() string {
	return strings.ToUpper(string(s))
}

func (s stripper) talking() string {
	return strings.ToUpper(string(s))
}

func (s stripper) commanding(i string) string {
	return "Command is: " + i
}
func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
			fmt.Println("我成功度过了panic")
		}
	}()
	var s stripper
	s = "zhangweilong"
	var t talKer = s
	fmt.Println(t.talking())
	fmt.Println(reflect.TypeOf(t))
	if c, ok := t.(stripper); ok {
		fmt.Println(reflect.TypeOf(c))
		fmt.Println(c)
		fmt.Println(c.commanding("fuck you"))
	}
	panic("I forgive my towel")
}

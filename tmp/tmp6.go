package main

import (
	"fmt"
	"strings"
)

type leaser int
type structTest struct {
	Name string
}

func (i leaser) talk() string {
	return strings.Repeat("pew ", int(i))
}
func (t structTest) talk() string {
	return strings.Repeat(t.Name, 3)
}

type temperature struct {
	leaser
	structTest
	// 这样写也可以
}

func main() {
	fmt.Println(temperature{leaser(3), structTest{"zhangweilong"}})
	t := temperature{leaser(10), structTest{Name: "liuwenzhuo"}}
	fmt.Println(t.structTest.talk())
	// 在结构嵌套的情况下，如果有多个子结构体都实现了同一个方法
	// 那么在调用的时候需要指定使用哪个子结构体的方法
	// 如果子结构体嵌入的时候指明了字段名，那么就用字段名进行调用，如果没有指明字段名
	// 就直接使用子结构体的名称进行调用，如上所示
	i := leaser(3)
	// leaser()是一个类型转换函数，将int类型的3转换成leaser类型的3
	fmt.Println(i.talk())
}

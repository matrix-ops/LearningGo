package main

import (
	"fmt"
	"reflect"
)

func parameter_add(prefix string, worlds ...string) {
	// 进行参数展开的形参会被当成一个切片处理
	fmt.Println(reflect.TypeOf(worlds))
	// []string
	for _, i := range worlds {
		fmt.Println(prefix + "" + i)
	}
}

func main() {
	parameter_add("New", "Mecury", "Venus", "Earth")
}

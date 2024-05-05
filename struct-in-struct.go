package main

import (
	"fmt"
)

type Car struct {
	Wheel Wheel
	Engine Engine
}

type Wheel struct {
	Size int
}

type Engine struct {
	Power int
	Type  string
}

func main() {
	c := Car{Wheel{Size: 10},Engine{Power: 10, Type: "V8"},}
	// 要么都写字段名，要么都不写
	fmt.Printf("%+v", c)
	//j, _ := json.MarshalIndent(c, "", "\t")
	//fmt.Println(string(j))

}

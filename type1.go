package main

import "fmt"

type getRowType func(row int) (string, string)

func drawTable(rows int, getRow getRowType) {

}

type kelvin float64

// type ADD func(a, b int) int
type sensor func(a, b kelvin) kelvin

// 可以函数也做成一种类，这样的话这个类的所有实例都是函数

func floatAdder(a, b kelvin) kelvin {
	return a + b
}
func main() {
	var s sensor = floatAdder
	// s是sensor类型的是一个实例
	fmt.Println(s(1.0, 2.0))
}

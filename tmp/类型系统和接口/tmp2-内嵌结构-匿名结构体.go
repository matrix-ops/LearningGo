package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int64
}
type Employee struct {
	Person
	// 此时Employee继承了Person内部所有字段
	// 如果在Employee这一层声明同名的字段，例如name，外部的name字段会覆盖掉内部的name字段
}

func main() {
	var f1 = new(Employee)
	fmt.Printf("%+v\n", f1)
	f1.name = "zhangweilong"
	f1.age = 25
	fmt.Println(f1)
}

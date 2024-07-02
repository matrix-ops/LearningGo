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
}

func main() {
	var f1 = new(Employee)
	fmt.Printf("%+v\n", f1)
	f1.name = "zhangweilong"
	f1.age = 25
	fmt.Println(f1)
}

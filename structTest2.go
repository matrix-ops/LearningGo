package main

import (
	fmt "fmt"
)

type Employee_T struct {
}

func (Employee_T) Strings() Employee_T {
	return Employee_T{}
}

func main() {
	var zhangweilong_V Employee_T
	fmt.Println(zhangweilong_V)
}

package main

import (
	"errors"
	"fmt"
)

var (
	errorTest = errors.New("这是一个巨大的错误")
)

func main() {
	fmt.Println(errorTest)
	fmt.Printf("%#v", errorTest)
}

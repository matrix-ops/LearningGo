package main

import "fmt"

func main() {
	city := "NYC"
	address := fmt.Sprintf("I live in %v", city)
	// Sprintf用于延迟打印
	//只有当Sprintf语句被赋值给变量，然后通过其他方法打印此变量的时候Sprintf函数的内容才会被打印
	fmt.Println(address)
}

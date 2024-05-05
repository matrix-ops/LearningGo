package main

import "fmt"

type date struct {
	Time int
}

func (d *date) updateDate() {
	s := fmt.Sprintf("I love computer")
	fmt.Println(s)
}

func main() {
	var someday *date
	fmt.Println(someday)
	someday.updateDate()
	// 方法可以接受空指针，只要不在方法内对空指针进行操作，就不会引发panic，
	// 但是如果接受了一个接受者而不针对它做任何操作，方法还有什么意义呢
}

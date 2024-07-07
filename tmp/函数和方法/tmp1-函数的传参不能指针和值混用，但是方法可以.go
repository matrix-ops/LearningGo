package main

import "fmt"

type person struct {
	name string
	sex  string
}

func fuck(p *person) {
	fmt.Println(p.name)
}
func main() {
	xiongmeiqi := person{name: "xiongmeiqi", sex: "women"}
	// fuck(xiongmeiqi)
	// 类型person不能用在*person的传参上
	fuck(&xiongmeiqi)
	// 这样是可以的
}

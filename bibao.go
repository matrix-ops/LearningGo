package main

import (
	"fmt"
	"time"
)

//func realFunc(x, y int) int {
//	x1 := x
//	y1 := y
//	fmt.Println(x1 , y1)
//	return x + y
//}

func main() {
	//h := realFunc(1,3)
	//fmt.Println(h)
	a := Fun()
	//d := Fun()
	b := a("hello ")
	c := a("hello ")
	c1 := a("mother fucker ")
	//e:=d("hello ")
	//f:=d("hello ")
	fmt.Println(b) //worldhello
	fmt.Println(c) //worldhello hello
	fmt.Println(c1)
	//fmt.Println(e)//worldhello
	//fmt.Println(f)//worldhello hello
}
func Fun() func(string) string {
	a := "world"
	return func(args string) string {
		// 保留了对父函数的引用，导致这个子函数在被返回之后，里面的值不会被销毁，一直常驻内存
		a += args
		// b := a
		// fmt.Println("a: " + a)
		// fmt.Println("b: " + b)
		time.Sleep(time.Second)
		return a
	}
}

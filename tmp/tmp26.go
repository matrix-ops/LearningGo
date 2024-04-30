package main

import (
	"fmt"
	"reflect"
	"time"
)

func gopher(id int, c chan int) {
	fmt.Println("id:", id)
	c <- id
}
func main() {
	var c = make(chan int)
	for i := 0; i < 5; i++ {
		go gopher(i, c)
	}
	for i := 0; i < 5; i++ {
		goChannel := <-c
		fmt.Println(goChannel)
		fmt.Println("gochannel的类型为：", reflect.TypeOf(goChannel))
	}
	time.Sleep(5 * time.Second)
}

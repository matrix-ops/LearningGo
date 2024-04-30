package main

import (
	"fmt"
	"reflect"
	"time"
)

func goRoutine(i int, c chan int) {
	c <- i
}
func main() {
	c := make(chan int)
	t1 := time.After(time.Second * 2)
	// time.After返回的是一个只读类型的chan，只能接收time.Time类型的值
	fmt.Println(reflect.TypeOf(t1))
	// <-chan time.Time
	go func() {
		fmt.Println(<-t1)
		fmt.Println("What")
	}()
	time.Sleep(time.Second * 3)
	go goRoutine(1, c)
	select {
	case t := <-time.After(time.Second * 2):
		// time.After语句返回的值是一个通道，Go Runtime里面的定时器会把时间值
		// 发送给time.After语句所返回的通道里面，t负责从此通道里面接受值
		//
		fmt.Println("t的值是：", t)
		// 此时t是一个time.Time类型
		// time.After()本身可以看做一个通道
	case i := <-c:
		fmt.Println("i的值是：", i)
	}
}

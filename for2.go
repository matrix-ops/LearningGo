// package main
//
// import (
// 	"fmt"
// 	"math/rand"
// )
//
// func main() {
// 	var degrees int = 0
// 	for {
// 		fmt.Println(degrees)
// 		degrees++
// 		if degrees >= 360 {
// 			degrees = 0
// 			if rand.Intn(2) == 0 {
// 				break
// 			}
// 		}
// 	}
// }
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var count int = 100
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		if rand.Intn(100) == 0 {
			break
		}
		count--
	}
	if count == 0 {
		fmt.Println("火箭发射！")
	} else {
		fmt.Println("程序错误，正在回滚")
	}
}

// 这个才是正确生成随机数的方法
// 火箭发射倒计时100秒，每次都有1/100的概率失败

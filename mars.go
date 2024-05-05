package main

import (
	"fmt"
)

func main() {
	var year int = 0
	var leap bool
	fmt.Println("输入一个年份：")
	fmt.Scanln(&year)
	leap = (year%4 == 0 && year%100 != 0) || year%400 == 0
	if leap {
		fmt.Println(year, "是一个闰年")
	} else {
		fmt.Println(year, "不是一个闰年")
	}
}

// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	var destination string
// 	fmt.Println("请输入一个目的地：")
// 	fmt.Scanln(&destination)
// 	if destination == "入口" {
// 		fmt.Println("入口这里真凉快")
// 	} else if destination == "洞穴" {
// 		fmt.Println("洞穴里面深不见底")
// 	} else if destination == "大山" {
// 		fmt.Println("好一座巍巍大山")
// 	} else {
// 		fmt.Println("我根本不知道你在说什么")
// 	}
// }
//
// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	var fruit1 = "apple"
// 	var fruit2 = "banana"
// 	if fruit1 > fruit2 {
// 		fmt.Println("True")
// 	} else {
// 		fmt.Println("False")
// 	}
// }
//
// package main
//
// import (
// 	"fmt"
// 	"strings"
// )
//
// func main() {
// 	fmt.Println("You find yourself in a dimly lit cavern.")
// 	var command = "walk outside"
// 	var exit = strings.Contains(command, "outside")
//
// 	fmt.Println("You leave the cave: ", exit)
// }

// package main
//
// import "fmt"
//
// func main() {
// 	const hoursPerDay = 24
// 	var days = 28
// 	var distance = 56000000 // km
// 	fmt.Println(distance/(days*hoursPerDay), "km/h")
// }

// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	var distance = 56000000 // km
// 	const hours = 672       // 28 * 24
// 	fmt.Println(distance/hours, "千米每小时")
// }

// ----
// package main
//
// import (
// 	"fmt"
// 	"math/rand"
// )
//
// func main() {
// 	var num = rand.Intn(10) + 1
// 	fmt.Println(num)
// 	num = rand.Intn(10) + 2
// 	fmt.Println(num)
// }

// -------------------------------
// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	const lightSpeed = 100800 // km/h
// 	var distance = 96300000   // km
// 	fmt.Println(distance/lightSpeed, "h")
// 	fmt.Println(distance/lightSpeed/24, "days")
// }

// ----------------------------
// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	fmt.Print("My weight on the surface of mars is ")
// 	fmt.Print(140 * 0.3783)
// 	fmt.Print(" kgs, and I would be ")
// 	fmt.Print(23 * 365 / 687)
// 	fmt.Print(" years old.\n")
// }

package main

import (
	"fmt"
)

func main() {
	var women string
	fmt.Println("请选择你的女人：")
	fmt.Scanln(&women)
	switch women {
	case "田淑清":
		fmt.Println("田淑清是你的老婆")
	case "黄淑婷":
		fmt.Println("黄淑婷是你的老婆")
	default:
		fmt.Println("黄宇婷")
	}
}

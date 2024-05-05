package main

import (
	"fmt"
)

func main() {
	var women string
	fmt.Println("请选择你的女人:")
	fmt.Scanln(&women)
	switch women {
	case "熊美琪":
		fmt.Println("熊美琪是你的老婆")
	case "黄淑婷":
		fmt.Println("黄淑婷不是你老婆")
		fallthrough
		// fallthrough语句会在本条分支执行完成之后
		// 无条件的执行下一个分支，在这个例子中，最后那个才是你老婆会被无条件的打印
	case "最后那个才是你老婆":
		fmt.Println("最后那个才是你老婆")
	default:
		fmt.Println("你没有老婆")
	}
}

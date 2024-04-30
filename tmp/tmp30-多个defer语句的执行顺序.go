package main

import "fmt"

func exampleFunction() int {
	defer fmt.Println("清理工作1")
	defer fmt.Println("清理工作2")

	if true {
		fmt.Println("提前返回")
		// defer语句只保证在return语句之前执行，这个Println语句会在defer语句之前执行
		return 1 // 第一个return语句
	}

	defer fmt.Println("不会执行的清理工作")

	fmt.Println("正常流程")
	return 0 // 第二个return语句，不会执行到
}

func main() {
	result := exampleFunction()
	fmt.Println("返回值是:", result)
}

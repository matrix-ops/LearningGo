package main

import (
	"fmt"
	"os"
)

func main() {
	characters, err := fmt.Fprintln(os.Stdout, "That may be a error")
	// 返回写入的字节数，和错误
	// Fprintln的第一个实参w必须是io.Writer类型，后面的实参用于写入w，它的返回值有两个
	// 一个是总共写入了多少个字节数，第二个是返回的错误
	// 因此编译运行此代码，屏幕上打印的那行输出是在上面那一行代码执行的时候打印的，而非在做错误处理的时候打印。
	fmt.Println(characters)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(1)
}

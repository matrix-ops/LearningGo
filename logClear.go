package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	appName := os.Getenv("APP_NAME")
	fmt.Println(appName)
	cmd := exec.Command("du", "/opt/log/app"+appName)
	err := cmd.Run()
	if err != nil {
		fmt.Println("你写的垃圾代码出问题了")
		os.Exit(127)
	}
	fmt.Println("命令执行完毕")
}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func main() {
	appName := os.Getenv("APP_NAME")
	dir := fmt.Sprintf("/opt/logs/app/%s", appName)
	if !exists(dir) {
		fmt.Println("日志路径不存在")
		os.MkdirAll(dir, os.ModePerm)
	}
	size, _ := DirSize(dir)
	// 80G
	if size > 85899345920 {
		fmt.Println("日志存储空间将满，即将执行清理，当前已用Bytes: ", size)
		cmd := exec.Command("/bin/rm", "-rf", dir)
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("清理完成")
		}
	} else {
		fmt.Println("日志存储空间未满,当前已用Bytes: ", size)
	}
}

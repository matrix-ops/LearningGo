package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.OpenFile("reqeust.logs.txt", os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("err:", err)
	}
	file, err := os.OpenFile("reqeust.logs.txt", os.O_APPEND, 0666)
	fmt.Println("err:", err, file)

}

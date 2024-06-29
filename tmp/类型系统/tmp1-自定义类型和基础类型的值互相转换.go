package main

import "fmt"

type Duration int64

func main() {
	var dur Duration
	dur = 1000
	fmt.Println(dur)
	// var dur1 Duration
	// dur1 = int64(1000)
	// 这样会报错
}

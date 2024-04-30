package main

import (
	"fmt"
	"time"
)

func main() {
	var t time.Time
	fmt.Println(t)
	// 0001-01-01 00:00:00 +0000 UTC
	fmt.Println(float64(t.YearDay()))
	// 1
	// 今年的第几天
	fmt.Println(t.Hour() / 24)
}

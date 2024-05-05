package main

import (
	"fmt"
	"math/rand"
	"time"
)

var year = 0000

func main() {
	rand.Seed(time.Now().UnixNano())
	leap := false
	startTime := "开始日期"
	count := 0
	for count < 10 {
		count++
		year = rand.Intn(9999) + 1
		month := rand.Intn(12) + 1
		dayInMonth := 31
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			leap = true
		}
		switch month {
		case 2:
			if leap == true {
				dayInMonth = 29
			} else {
				dayInMonth = 28
			}
		case 4, 6, 9, 11:
			dayInMonth = 30
		}
		day := rand.Intn(dayInMonth) + 1
		fmt.Println(startTime, ":", year, month, day)
	}
}

// package main
//
// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )
//
// var era = "AD"
//
// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	year := 2020
// 	switch month := rand.Intn(12) + 1; month {
// 	case 2:
// 		day := rand.Intn(28) + 1
// 		fmt.Println(era, year, month, day)
// 	case 4, 6, 9, 11:
// 		day := rand.Intn(30) + 1
// 		fmt.Println(era, year, month, day)
// 	default:
// 		day := rand.Intn(31) + 1
// 		fmt.Println(era, year, month, day)
// 	}
// }

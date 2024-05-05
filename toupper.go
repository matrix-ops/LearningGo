package main

import "fmt"

type huashi float64
type Celsius float64

func FtoC(f huashi) Celsius {
	c := f + 32.0
	return Celsius(c)
}

func main() {
	var c float64 = 0.0
	fmt.Println(FtoC(c))
}

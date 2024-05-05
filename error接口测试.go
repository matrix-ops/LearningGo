package main

import (
	"fmt"
	"math"
)

type dualError struct {
	Num     float64
	problem string
}

func (e dualError) Error() string {
	return fmt.Sprintf("Wrong!!!,because \"%f\" is a negative number, problem: \"%v\"", e.Num, e.problem)
}
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, dualError{Num: f, problem: "You are so stupid!"}
	}
	return math.Sqrt(f), nil
}
func main() {
	result, err := Sqrt(-13)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

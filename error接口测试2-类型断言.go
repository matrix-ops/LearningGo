package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

const rows, columns = 9, 9

type Grid [rows][columns]int8
type SudokuError []error

// Error 方法返回一个或多个以逗号分隔的错误
// SudokuError类型具有了Error方法，并且方法签名满足error接口的要求之后
// SudokuError本身也满足了error接口的要求，它自身就是一个error接口类型
func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
		// 只有使用了Error方法,err中的内容才会变成字符串,
	}
	return strings.Join(s, ", ")
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}
func (g *Grid) Set(row, column int, digit int8) error {
	var errs SudokuError
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
		// 只要一个类型具有Error方法，它就能被当做error返回
	}
	g[row][column] = digit
	return nil
}

func main() {
	var g Grid
	err := g.Set(10, 0, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\nhaha\n", ErrBounds)
	fmt.Println(ErrBounds.Error())
	if errs, ok := err.(SudokuError); ok {
		// 如果err变量是SudokuError类型，那么它作为SudokuError类型赋值给errs
		fmt.Println(len(errs))
		for _, err := range errs {
			fmt.Println(err)
		}
	}
}

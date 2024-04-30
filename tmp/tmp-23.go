package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigits = errors.New("invalid digit")
)

// 这里是在var里面所以用不了:=

const rows, columns = 9, 9

type Grid [rows][columns]int

// 一个二维数组

func (g *Grid) Set(r, c int, digit int) error {
	if !inBound(r, c) {
		return errors.New("out of bounds")
	}
	(*g)[r][c] = digit
	// 注意，正常来说，由于接受者是一个指针类型，那么调用的时候也只能由指针类型的变量来进行调用方法
	// 但是Go语言提供了自动取内存地址的人体工程学功能，当一个类型的指针类型具有某类方法的时候
	// 使用这个类型的值变量去调用给指针接受者设置的方法的情况下，会自动去值变量的内存地址，以指针的形式去调用它
	return nil
}
func inBound(r, c int) bool {
	if r < 0 || r >= rows {
		return false
	}
	if c < 0 || c >= columns {
		return false
	}
	return true
}
func main() {
	var g Grid
	err := g.Set(8, 8, 9)
	if err != nil {
		fmt.Println("An error occurred: ", err)
		os.Exit(1)
	}
}

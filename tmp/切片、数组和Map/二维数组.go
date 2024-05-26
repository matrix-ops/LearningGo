package main

import "fmt"

func main() {
	var board [8][8]string
	// 声明一个二维数组
	board[0][0] = "r"
	board[0][7] = "r"
	fmt.Println("board[0]的长度是：", len(board[0]), ", board[0]的容量是：", cap(board[0]))
	// 8,8
	for column := range board[1] {
		board[1][column] = "p"
	}
	// 把board这个二维数组的第二个元素赋值为一个总共八个元素，每个元素的值都是"p"的数组
	fmt.Println(board)
	var b string = board[0][0]
	fmt.Println(b)
	// r
	// 数组的元素可以直接当成值来使用
}

package main

import (
	"fmt"
	"sort"
)

// 第一个参数定义一个字符串切片，第二个函数定义了一个函数原型，用于决定排序的方向，
// 使用sortStrings函数时可以直接以匿名函数的方式给出第二个参数
func sortStrings(s []string, more func(i, j int) bool) {
	if more == nil {
		more = func(i, j int) bool {
			return s[i] > s[j]
			// 这里的大小于符号用于决定排序的方向
			// 可以根据不同的需求来变动此处的代码，例如用len()函数根据字符串的长度来进行排序
		}
	}
	sort.Slice(s, more)
}
func main() {
	xingxing := []string{"onion", "Alpha", "Carrot", "celery"}
	sortStrings(xingxing, func(i, j int) bool { return xingxing[i] < xingxing[j] })
	fmt.Println(xingxing)
}

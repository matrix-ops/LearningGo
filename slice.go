package main

import (
	"fmt"
	"reflect"
	"sort"
)

func PrintNewNumber(numbers ...int) []int {
	// 用这种展开型的写法来写形参，传参传进来以后会被存储到一个同名的变量里面
	fmt.Println("numbers的值是 ", numbers)
	fmt.Println("numbers的类型是", reflect.TypeOf(numbers))
	// reflect.Typeof用来确定一个变量的类型
	newNumbers := make([]int, len(numbers))
	for i := range numbers {
		newNumbers[i] = numbers[i] + 100
	}
	return newNumbers
}

func main() {
	intList := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(PrintNewNumber(intList...))
	stringSlice := make([]string, 10)
	stringSlice[0] = "Test"
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	fmt.Println(intList)
	fmt.Println(cap(stringSlice), len(stringSlice))
}

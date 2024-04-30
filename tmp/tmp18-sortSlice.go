package main

import (
	"fmt"
	"sort"
)

func main() {
	s2 := []string{
		"abc", "bac", "cab",
	}
	s3 := []string{
		"cbdad",
		"a",
		"abc",
	}
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	sort.Slice(s3, func(i, j int) bool {
		return len(s3[i]) < len(s3[j])
	})
	// sort.Slice在第二个匿名函数的返回值为true的时候，按照函数体内部定义的逻辑进行排序
	fmt.Println(s2)
	// [abc bac cab]
	fmt.Println(s3)
	// [a abc cbdad]
}

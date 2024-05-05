package main

import "fmt"

func main() {
	slicE := []int{1, 2, 3, 4, 5}
	groups := make(map[string][]int, 10)
	homes := make(map[string][6]int, 5)
	var isPresent bool
	groups["Money1"] = []int{1, 2, 3, 4, 5, 6}
	// 映射的值可以是切片
	groups["Money2"] = []int{1, 2, 3, 4, 5, 6}
	groups["Money3"] = []int{1, 2, 3, 4, 5, 6}
	homes["Money1"] = [6]int{1, 2, 3, 4, 5, 6}
	// 映射的值可以是数组
	// if _, isPresent = groups["Money4"]; isPresent {
	// 	fmt.Println(groups)
	// }
	if _, isPresent = homes["Money1"]; isPresent {
		fmt.Println("Map homes: ")
		fmt.Println(homes)
	}
	// OK语法，如果groups中存在Money4键位，那么isPresent为真
	// 如果不需要使用第一个值，可以留空
	for i := range groups {
		fmt.Println("groups索引值: " + i)
		// 如果只写一个值，遍历的将是键
	}
	for j := range slicE {
		fmt.Println(j)
	}
	if _, ok := groups["Moon"]; ok {
		fmt.Println("Key Moon Exist")
	} else {
		fmt.Println("Key Moon no Exist")
	}
	map2 :=

}

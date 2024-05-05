// 有3,4,5,6四个数字，统计他们能组成多少个互不相同且无重复字数的三个数
// 将百位数字相同的三位数同一行输出
package main

import "fmt"

func main() {
	var numbers = [4]int{3, 4, 5, 6}
	end := 0
	for _, i := range numbers {
		for _, j := range numbers {
			for _, k := range numbers {
				if i != j && j != k && i != k {
					end = i*100 + j*10 + k
					fmt.Println(end)
				}
			}
		}
	}
}

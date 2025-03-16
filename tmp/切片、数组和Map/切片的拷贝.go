package main

import "fmt"

func main() {
	s1 := []string{"Apple", "Orange", "Banana", "Bear"}
	fmt.Println("s1一开始是：", s1)
	fmt.Println("现在开始测试拷贝, 新建切片s2，复制s1到s2，然后更改s1")
	s2 := s1
	s1[0] = "strawberry"
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	fmt.Println("现在开始测试深拷贝,新建s3并初始化，然后更改s1")
	s3 := make([]string, len(s1))
	copy(s3, s1)
	s1[0] = "watermelon"
	fmt.Println("s1: ", s1)
	fmt.Println("s3: ", s3)
}

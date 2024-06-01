package main

import "fmt"

func checkSliceNil(s []int) {
	if s == nil {
		fmt.Println("slice is nil!")
	} else {
		fmt.Println("slice is not nil!")
	}

}
func main() {
	s1 := make([]int, 3, 10)
	s1 = []int{1, 2, 3}
	fmt.Println(s1, cap(s1), len(s1))
	// 使用make初始化切片之后重新赋值一个字面量
	// 切片的容量和长度都会变成字面量底层数组的容量和长度
	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(s2, cap(s2), len(s2))
	// 10, 10

	s4 := s2[4:8]
	// 索引最大到9
	fmt.Println(s4, "s4的容量是：", cap(s4), "s4的长度是：", len(s4))
	// 注意，在切片上的基础上新增切片的时候，计算容量应该按照右边的没有用的数组元素来计算
	// 也就是说默认的第三个参数其实是最大的那个索引数字,在这个例子中就是[4:8:8]
	// 第三个参数指的是新切片的容量，从第一个索引计算，到到哪个索引为止，这一段切出来的数组作为底层数组
	// 在这个例子中，就是从索引4开始到索引9这一段的空间作为底层数组
	// 也就是原数组的2,3,4,5,6
	// 遵循左闭右开区间原则
	s3 := s2[2:7:7]
	fmt.Println(s3, cap(s3), len(s3))
	// cap:5 len:5
	s5 := s2[2:7]
	fmt.Println(s5, cap(s5), len(s5))
	// cap:8 len:5
	// cap从2-9

	checkSliceNil(s1)
	// false
}

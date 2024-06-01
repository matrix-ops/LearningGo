package main

import "fmt"

func main() {
	// 如果新切片的长度和容量不一致，在持续对新切片执行append操作时，容易搞不清楚此次append操作是否会影响原有的底层数组
	// 还是新建一个底层数组
	// 因此，将新切片的长度和容量都设置相同是一个好习惯，这样可以确保新切片的第一个append操作就是新建了一个底层数组
	s1 := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	fmt.Println("s1的长度是：", len(s1), "s1的容量是：", cap(s1))
	// 5, 5
	s2 := s1[2:3:3]
	s2 = append(s2, "Kiwi")
	fmt.Println(s2)
	// [Plum, Kiwi]
	fmt.Println("s2的长度是：", len(s2), "s2的容量是：", cap(s2))
	// 2 , 2
	s2 = append(s2, s1...)
	// 切片元素展开，这个操作会把s1的所有元素都追加到s2
	// 这个操作并不会导致s2的容量变成追加完成之后的长度7的两倍
	// 在底层数组的容量翻倍仍然无法满足切片的长度需求时，容量将被设置成和长度相同
	// 在这个例子中容量翻倍2*2变成4以后仍然无法满足追加的需求，因此容量将被设置成7
	fmt.Println("s2的长度是：", len(s2), "s2的容量是：", cap(s2))
	// 7, 7
	s2 = append(s2, "Pear")
	fmt.Println("s2的长度是：", len(s2), "s2的容量是：", cap(s2))
	// 8 , 14
}

package main

import "fmt"

func Counter(element ...string) {
	// 第二种用法：可变长度参数
	// 可以给函数定义一个任意长度的参数,这样调用方可以传递任意个参数给函数
	// 接收到的参数将被存储为一个切片
	for _, i := range element {
		fmt.Println(i)
	}
	fmt.Println("element的长度是：", len(element))
}

func main() {
	// 第一种用法：让编译器自己根据复合字面量来计算数组的长度和容量
	a1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(a1), cap(a1))
	// 5, 5
	// 第二种用法
	Counter("zhangweilong", "Future", "Dream")
	// 第三种用法，参数展开
	s1 := []int{1, 2, 3, 4, 5}
	var s2 []int
	s2 = append(s2, s1...)
	fmt.Println("s2的内容是：", s2)
	// 1,2,3,4,5
	// 第三种用法的介绍就到这里，下面是关于append函数在对容量为0的切片执行append操作的时候，一些验证
	//
	// 结论是append函数并不是简单的在底层数组容量不够用的情况下，就执行新建一个容量两倍于原底层数组的新数组
	// 在这中间有一个内存对齐的操作，因此在对容量为0的切片执行append操作时，会出现一些出人意料的结果
	// 例如追加一个五元素的切片到容量为0的切片时，新切片的底层数组容量是6
	// 追加追加一个六元素的切片到容量为0的切片时，新切片的底层数组容量是8
	// 在常规追加操作时（指的是原切片的容量不是0的时候）,也只有在一定容量以内的切片才会在新增的时候得到一个两倍容量的数组
	// Go 1.18之前，原slice容量小于1024的时候，新的slice的容量将会是原slice的两倍，
	// 大于1024的时候, 新slice的容量近似于原slice的1.25倍
	// Go 1.18之后，原slice容量小于256的时候，新slice的容量将会是原slice的两倍，
	// 大于256的时候，新slice的容量近似于(oldcap+(oldcap+3*256)/4)的策略
	// 之所以两者都说是近似，就是因为上面提到的内存对齐原则的存在，growslice函数在计算完新slice的容量之后，会使用
	// roundupsize函数进行内存对齐一次

	//

	fmt.Println("s2的长度是：", len(s2), "s2的容量是：", cap(s2))
	// 5, 6
	s2 = append(s2, 6)
	fmt.Println("s2的长度是：", len(s2), "s2的容量是：", cap(s2))
	fmt.Println(s2)

	var s3 []int
	fmt.Println(cap(s3))
	s3 = append(s3, s1...)
	fmt.Println("s3的长度是：", len(s3), "s3的容量是：", cap(s3))

	s4 := make([]int, 6)
	// make函数的第二个参数默认是长度，第三个参数才是切片的容量
	s4 = append(s4, s1...)
	fmt.Println("s4的长度是：", len(s4), "s4的容量是：", cap(s4))

	s5 := make([]int, 7, 7)
	s5 = append(s5, s1...)
	fmt.Println("s5的长度是：", len(s5), "s5的容量是：", cap(s5))
	var s6 []int
	s6 = append(s6, 1)
	fmt.Println("s6的长度是：", len(s6), "s6的容量是：", cap(s6))
	s6 = append(s6, 2)
	fmt.Println("s6的长度是：", len(s6), "s6的容量是：", cap(s6))
	s6 = append(s6, 3)
	fmt.Println("s6的长度是：", len(s6), "s6的容量是：", cap(s6))
	s6 = append(s6, 4)
	fmt.Println("s6的长度是：", len(s6), "s6的容量是：", cap(s6))
	s6 = append(s6, 5)
	fmt.Println("s6的长度是：", len(s6), "s6的容量是：", cap(s6))
	// 5, 8
}

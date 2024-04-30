package main

// 指针的作用就是指向
// 指向另外一个东西的内存地址

import "fmt"

func main() {
	name := "zhangweilong"
	sun := "solar"
	people1 := &name
	people2 := people1
	fmt.Println(people1, people2)
	// 两者指向相同的内存地址，因此输出完全相同
	people2 = &sun
	fmt.Println(people1, people2)
	// 指向已经不一样
	fmt.Println(*people1, *people2)
	sun = "in my dick"
	fmt.Println(*people1, *people2)
}

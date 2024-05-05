package main

import "fmt"

func main() {
	plantes := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter"}
	plantes_plus := append(plantes, "Saturn")
	// 往底层数组初始容量只有5的切片新增数据，Go将重新新建一个容量为原来底层数组两倍的新底层数组
	fmt.Println(plantes_plus)
	fmt.Println(cap(plantes), cap(plantes_plus))
	fmt.Printf("%T\n", plantes)
	fmt.Printf("%T\n", plantes_plus)
	// 打印plantes的类型
}

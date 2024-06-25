package main

import "fmt"

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// 等待所有协程完成后再退出
	for _ = range values {
		<-done
	}
	var mapreduce = make(map[string]int)
	// fmt.Println(cap(mapreduce))
	// cap函数不能用于map类型
	fmt.Println(len(mapreduce))
	// len函数可以
}

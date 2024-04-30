package main

import (
	"fmt"
	"sync"
)

type Visited32 struct {
	mu      sync.Mutex
	visited map[string]int
}

func (v *Visited32) VisitedLink(url string) int {
	// 因为这里使用了指针类型来作为接收者
	// 那么不管实参是指针版本还是非指针版本，这里实际上都会当成指针处理
	// 这是Go语言的人体工程学设计
	v.mu.Lock()
	defer v.mu.Unlock()
	// 如果不适用互斥锁，在goroutine中使用这个方法将会报错
	count := v.visited[url]
	count = count + 1
	v.visited[url] = count
	return count
}

func GoRoutineVisited(v Visited32) {
	for _, i := range []string{"Google", "Baidu", "Bing"} {
		v.VisitedLink(i)
	}
}

func main() {
	v := Visited32{visited: map[string]int{}}
	v.visited["Google"] = 1
	for i := 0; i < 10; i++ {
		go GoRoutineVisited(v)
	}
	fmt.Println(v)
}

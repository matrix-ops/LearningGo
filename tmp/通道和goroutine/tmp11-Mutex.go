package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.Mutex
	accessRecord := make(map[string]int)
	accessRecord["baidu"] = 1
	mutex.Lock()
	{
		count := accessRecord["baidu"]
		count++
		accessRecord["baidu"] = count
	}
	mutex.Unlock()
	fmt.Println("accessRecord:", accessRecord)
}

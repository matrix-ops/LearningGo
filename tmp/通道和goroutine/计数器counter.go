package main

import (
	"fmt"
	"sync"
)

package main

import (
"fmt"
"sync"
)

var (
	counter int
	mutex   sync.Mutex
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	go intCounter(1)
	go intCounter(2)
	wg.Wait()
	fmt.Println(counter)
}

func intCounter(i int) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		mutex.Lock()
		{
			value := counter
			value += 1
			counter = value
		}
		mutex.Unlock()
	}
}

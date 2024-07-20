package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			for i := 'A'; i < 'A'+26; i++ {
				fmt.Printf("%c  ", i)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c  ", char)
			}
		}
	}()
	fmt.Println("Waiting to finish.")
	wg.Wait()
	fmt.Println("Terminating Program.")
}

package main

import "fmt"

func squares(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i <= n; i++ {
			ch <- i * i
		}
		close(ch)
	}()
	return ch
}

func main() {
	term := squares(5)
	fmt.Println(<-term)
	fmt.Println(<-term)
	fmt.Println(<-term)
	fmt.Println(<-term)
	fmt.Println(<-term)
}

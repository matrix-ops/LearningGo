package main

import (
	"fmt"
	"strings"
	"time"
)

func sourceGopher(downstream chan string) {
	for _, v := range []string{"Apple", "Orange", "Banana"} {
		downstream <- v
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

func printGopher(downstream chan string) {
	for v := range downstream {
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	go printGopher(c1)
	time.Sleep(time.Second * 1)
}

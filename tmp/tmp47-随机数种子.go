package main

import (
	"fmt"
	"math/rand"
)

func init() {
}

func main() {
	rand.New(rand.NewSource(32))
	fmt.Println(rand.Intn(100))
	rand.New(rand.NewSource(32))
	fmt.Println(rand.Intn(100))
	rand.New(rand.NewSource(32))
	fmt.Println(rand.Intn(10))
}

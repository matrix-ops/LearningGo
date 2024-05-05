package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	future := time.Unix(22622780801, 0)
	countdown, err := strconv.Atoi("100")
	fmt.Println(future)
	fmt.Printf("%T\n", future)
	fmt.Println(countdown, err)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "I Love You"
	test := strings.Contains(data, "Love You")
	fmt.Println(test)
}

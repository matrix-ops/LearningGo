package main

import (
	"fmt"
	"strings"
)

func main() {
	message := "L fdph, L vdz, L frqtxhuhg."
	for i := 0; i < len(message); i++ {
		var character1 = message[i]
		if (character1 >= 'A' && character1 <= 'Z') || (character1 >= 'a' && character1 <= 'z') {
			if strings.Contains("abcABC", string(character1)) {
				character1 = character1 + 23
				fmt.Printf("%c", character1)
			} else {
				character1 = character1 - 3
				fmt.Printf("%c", character1)
			}
		} else {
			fmt.Printf("%c", character1)
		}
	}
}

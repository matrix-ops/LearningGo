package main

import (
	"fmt"
	"strings"
)

func hyperSpace(worlds []string) {
	for i, _ := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
		worlds[i] = "New_" + worlds[i]

	}
}

func main() {
	planets := []string{" Mercury ", "Venus ", "Earth", "Mars", "Jupiter"}
	hyperSpace(planets)
	fmt.Println(strings.Join(planets, " "))
	fmt.Println(planets)
	fmt.Printf("%T", planets)
}

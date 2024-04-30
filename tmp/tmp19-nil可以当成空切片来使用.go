package main

import "fmt"

func main() {
	soup := tang(nil)
	fmt.Println(soup)
}
func tang(slice []string) []string {
	return append(slice, "bitch", "you", "mother")
}

package main

import "fmt"

func main() {
	message := "uv vagreangvbany fcnpr fgngvba"
	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
				// 例如如果是第14个字母，14+13=27，字母表上没有第27个字母，于是减去26之后就变成了第一个字母
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()
}

package main

import "fmt"

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	keyIndex := 0
	var message uint8
	for i, j := 0, 0; i < len(cipherText); i++ {
		cipherCharacter := cipherText[i]
		currentKeyword := keyword[j]
		decipher := currentKeyword - cipherCharacter
		if decipher < 65 {
			message = cipherCharacter + (26 - (currentKeyword - 65))
		} else {
			message = cipherCharacter - (currentKeyword - 65)
		}
		fmt.Printf("%c", message)
		keyIndex++
		keyIndex %= len(keyword) - 1
	}
}

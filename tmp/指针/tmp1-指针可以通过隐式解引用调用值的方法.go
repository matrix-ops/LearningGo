package main

import "fmt"

type Animal interface {
	Speak() string
}
type Dog struct {
}

func (d Dog) Speak() string {
	return "I am a dog"
}

func main() {
	var dog *Dog = &Dog{}
	fmt.Println(dog.Speak())
}

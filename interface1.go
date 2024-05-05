package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type locationL struct {
	sol int
}

type lawer int

func (l locationL) talk() string{
	return "HaHaHaHa"
}

func(l lawer) talk() string {
	return strings.Repeat("Fuck you", 3)
}

func main()  {
	var l1 talker
	l1 = new(locationL)
	l2 := locationL{sol: 15}
	fmt.Println(l1.talk())
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Printf("%+v", l1)
	//l = lawer(3)
	//fmt.Println(l.talk())
}

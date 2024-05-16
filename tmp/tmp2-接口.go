package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

func (t talker) talkShout(s string) {
	t.talk()
	// 错误的，接口自身不可以有方法
}

type laser int

func (l laser) shout(s string) {
	fmt.Println(strings.Repeat(s, int(l)))
	// strings.Repeat第一个参数类型是string，第二个参数类型是int
}

func (l laser) talk() string {
	return "son of bitch"
}
func main() {
	var l laser = 3
	l.shout("ha")
	// hahaha
	var s talker = l
	s.talk()
	// son of bitch
}

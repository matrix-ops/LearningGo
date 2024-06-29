package main

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}
type admin struct {
	person user
	level  int
}

func main() {
	zhangweilong := admin{
		person: user{
			name:       "zhangweilong",
			email:      "vhuumk@gmail.com",
			ext:        123,
			privileged: true,
		},
		level: 123,
	}
	fmt.Printf("%+v\n", zhangweilong)
}

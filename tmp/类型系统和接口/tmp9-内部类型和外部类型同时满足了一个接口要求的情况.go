package main

import "fmt"

type notifier interface {
	notify()
}
type admin struct {
	user
	level int
}
type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}
func sendNotifyEmail(n notifier) {
	n.notify()
}

func main() {
	ad := admin{
		user: user{
			name:  "zhangweilong",
			email: "vhuumk@gmail.com",
		},
		level: 1,
	}
	ad.notify()
	ad.user.notify()
	sendNotifyEmail(&ad.user)
	// 调用内部类型user
	sendNotifyEmail(&ad)
	// 调用外部类型admin
}

package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) sendmain() {
	fmt.Printf("Sending mail to %s via %s \n", u.name, u.email)
}
func (u *user) changemail(email string) {
	u.email = email
}
func main() {
	u := user{"Bob", "bob@gmail.com"}
	u.sendmain()
	u.changemail("Bob@gmail.com")
	fmt.Println(u)
}

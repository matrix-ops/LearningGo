package main

import (
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"os"
)

func main() {
	e := email.NewEmail()
	e.From = "zhangweilong <zhangweilong@payeco.com>"
	e.To = []string{"xu.hua@payeco.com"}
	e.Subject = "test"
	e.Text = []byte("test")
	err := e.Send("mail.payeco.com:465", smtp.PlainAuth("zhangweilong", "zhangweilong@payeco.com", "8h4k3d7n", "mail.payeco.com"))
	if err != nil {
		log.Fatal(err)
	} else {
		os.Exit(1)
	}
}

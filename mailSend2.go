package main

import (
	"crypto/tls"
	"fmt"
	mail "gopkg.in/gomail.v2"
)

func main() {
	msg := mail.NewMessage()
	// 这是一个结构体指针
	// 发件人
	msg.SetHeader("From", "zhang.weilong@payeco.com")
	// 收件人
	msg.SetHeader("To", "zhang.weilong@payeco.com")
	// 设置邮件标题
	msg.SetHeader("Subject", "Go 语言发送邮件")
	// 内容
	msg.SetBody("text/html", "<h3>测试邮件发送")
	fmt.Printf("%T", msg)
	dialer := mail.NewDialer("mail.payeco.com", 25, "zhang.weilong@payeco.com", "8h4k3d7n")
	// 不校验证书
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	fmt.Printf("%+v", dialer)
	if err := dialer.DialAndSend(msg); err != nil {
		panic(err)
	}
}

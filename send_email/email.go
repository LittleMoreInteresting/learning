package main

import (
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(userName, authCode, host, portStr, mailTo, sendName string, subject, body string) error {
	port, _ := strconv.Atoi(portStr)
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, sendName))
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, port, userName, authCode)
	err := d.DialAndSend(m)
	return err
}

func main() {

	var (
		host       = "smtp.126.com"
		port       = "25"
		senderName = "Blog"
		user       = "devin1015@126.com"
		password   = "****"
	)
	err := SendMail(user, password, host, port, "1208818417@qq.com", senderName, "验证码：", "1212")
	panic(err)
}

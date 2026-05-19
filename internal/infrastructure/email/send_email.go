package email

import (
	"fmt"
	"net/smtp"
)

type SmtpClient struct {
	addr string
	auth smtp.Auth
}

func CreateSmtpClient(addr, host, passwd string) *SmtpClient {
	return &SmtpClient{
		addr: addr,
		auth: smtp.PlainAuth("", addr, passwd, host),
	}
}

func (c *SmtpClient) Send(smtpAddr, headline, body string) {
	msg := "From: " + c.addr + "\n" +
		"To: " + c.addr + "\n" +
		"Subject: " + headline + "\n\n" +
		body
	fmt.Println("addr", smtpAddr)
	err := smtp.SendMail(smtpAddr, c.auth, c.addr, []string{c.addr}, []byte(msg))
	if err != nil {
		fmt.Println("failed to send mail", err)
	}
}

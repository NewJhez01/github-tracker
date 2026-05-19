package email

import (
	"fmt"
	"net/smtp"
)

type SmtpClient struct {
	addr     string
	smtpAddr string
	auth     smtp.Auth
}

func NewSmtpClient(addr, host, passwd, smtpAddr string) *SmtpClient {
	return &SmtpClient{
		addr:     addr,
		auth:     smtp.PlainAuth("", addr, passwd, host),
		smtpAddr: smtpAddr,
	}
}

func (c *SmtpClient) Send(headline, body string) {
	msg := "From: " + c.addr + "\n" +
		"To: " + c.addr + "\n" +
		"Subject: " + headline + "\n\n" +
		body
	err := smtp.SendMail(c.smtpAddr, c.auth, c.addr, []string{c.addr}, []byte(msg))
	if err != nil {
		fmt.Println("failed to send mail", err)
	}
}

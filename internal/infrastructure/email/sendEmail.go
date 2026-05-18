package email

import (
	"fmt"
	"net/smtp"
)

type SmtpClient struct {
	addr string
	auth smtp.Auth
}

func CreateSmtpClient(addr string) *SmtpClient {
	return &SmtpClient{
		addr: addr,
		// to do handle secrets
		auth: smtp.PlainAuth("", addr, "************", "smtp.gmail.com"),
	}
}

func (c *SmtpClient) Send(headline, body string) {
	msg := "From: " + c.addr + "\n" +
		"To: " + c.addr + "\n" +
		"Subject: " + headline + "\n\n" +
		body
	err := smtp.SendMail("smtp.gmail.com:587", c.auth, c.addr, []string{c.addr}, []byte(msg))
	if err != nil {
		fmt.Println("failed to send mail", err)
	}
}

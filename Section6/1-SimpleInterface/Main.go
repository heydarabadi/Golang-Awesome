package main

import "fmt"

type ISendNotification interface {
	Send()
}

type Mail struct {
	Message string
	Subject string
}

type WhatsApp struct {
	Message string
	Subject string
}

func (mail *Mail) Send() {
	fmt.Printf("%s\n", mail.Message)
}

func (whatsApp *WhatsApp) Send() {
	fmt.Printf("%s\n", whatsApp.Message)
}

func main() {
	mail := Mail{
		Message: "Hello World",
		Subject: "Main",
	}

	mail.Send()

	whatsApp := WhatsApp{
		Message: "Hello World",
		Subject: "WhatsApp",
	}

	whatsApp.Send()
}

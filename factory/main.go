// Factory centraliza la creación de objetos que comparten una interfaz.
package main

import "fmt"

type Notification interface {
	Send(message string)
}

type EmailNotification struct{}

func (EmailNotification) Send(message string) {
	fmt.Println("Email:", message)
}

type SMSNotification struct{}

func (SMSNotification) Send(message string) {
	fmt.Println("SMS:", message)
}

// NewNotification decide qué implementación crear según el canal pedido.
func NewNotification(channel string) Notification {
	if channel == "sms" {
		return SMSNotification{}
	}
	return EmailNotification{}
}

func main() {
	notification := NewNotification("email")
	notification.Send("La orden fue confirmada")
}

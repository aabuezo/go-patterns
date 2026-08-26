// DIP - Principio de inversión de dependencias
package main

import "fmt"

// El DIP establece que los módulos de alto nivel no deberían depender de
// módulos de bajo nivel, sino de abstracciones.
// Ambos deberían depender de abstracciones (interfaces), no de implementaciones concretas.

// MessageSender es la abstracción que necesita el módulo de alto nivel.
// No importa si el mensaje se envía por email, SMS o cualquier otro medio.
type MessageSender interface {
	Send(message string)
}

// EmailSender es un módulo de bajo nivel: conoce los detalles de enviar
// mensajes por email y cumple con la interfaz MessageSender.
type EmailSender struct{}

func (EmailSender) Send(message string) {
	fmt.Println("Enviando email:", message)
}

// OrderService es un módulo de alto nivel porque contiene una regla del
// negocio: confirmar una orden.
// Depende de la abstracción MessageSender, no directamente de EmailSender.
type OrderService struct {
	sender MessageSender
}

// La dependencia se recibe desde afuera. Esto se llama inyección de dependencias.
func NewOrderService(sender MessageSender) OrderService {
	return OrderService{sender: sender}
}

func (o OrderService) ConfirmOrder() {
	// OrderService solo sabe que puede enviar un mensaje.
	// No necesita conocer los detalles del email.
	o.sender.Send("Tu orden fue confirmada")
}

func main() {
	// EmailSender puede reemplazarse por cualquier otro tipo que implemente
	// MessageSender, sin modificar OrderService.
	emailSender := EmailSender{}
	orderService := NewOrderService(emailSender)

	orderService.ConfirmOrder()
}

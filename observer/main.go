// Observer notifica a varios objetos cuando cambia un estado.
package main

import "fmt"

type Observer interface {
	Update(product string)
}

type EmailSubscriber struct {
	address string
}

func (s EmailSubscriber) Update(product string) {
	fmt.Println("Email a", s.address, ":", product, "disponible")
}

type ProductStock struct {
	observers []Observer
}

func (s *ProductStock) Subscribe(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s ProductStock) AddProduct(product string) {
	for _, observer := range s.observers {
		observer.Update(product)
	}
}

func main() {
	stock := ProductStock{}
	stock.Subscribe(EmailSubscriber{address: "ana@example.com"})
	stock.Subscribe(EmailSubscriber{address: "juan@example.com"})
	stock.AddProduct("auriculares")
}

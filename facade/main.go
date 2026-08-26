// Facade ofrece una operación simple sobre un sistema con varios pasos.
package main

import "fmt"

type Inventory struct{}

func (Inventory) Reserve(product string) bool {
	fmt.Println("Reservando", product)
	return true
}

type Payment struct{}

func (Payment) Charge(amount float64) {
	fmt.Println("Cobrando", amount)
}

type Shipping struct{}

func (Shipping) Schedule(product string) {
	fmt.Println("Coordinando envío de", product)
}

// OrderFacade oculta los pasos necesarios para crear una orden.
type OrderFacade struct {
	inventory Inventory
	payment   Payment
	shipping  Shipping
}

func (f OrderFacade) PlaceOrder(product string, amount float64) {
	if f.inventory.Reserve(product) {
		f.payment.Charge(amount)
		f.shipping.Schedule(product)
		fmt.Println("Orden creada")
	}
}

func main() {
	facade := OrderFacade{Inventory{}, Payment{}, Shipping{}}
	facade.PlaceOrder("teclado", 80)
}

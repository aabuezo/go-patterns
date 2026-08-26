// OCP - Principio abierto/cerrado
package main

import "fmt"

// Discount representa una forma de calcular un descuento.
type Discount interface {
	Apply(price float64) float64
}

// NoDiscount no modifica el precio.
type NoDiscount struct{}

func (NoDiscount) Apply(price float64) float64 {
	return price
}

// TenPercentDiscount aplica un descuento del 10%.
type TenPercentDiscount struct{}

func (TenPercentDiscount) Apply(price float64) float64 {
	return price * 0.90
}

// Checkout calcula el precio usando cualquier descuento que cumpla la
// interfaz. No necesita conocer cada tipo de descuento.
type Checkout struct{}

func (Checkout) FinalPrice(price float64, discount Discount) float64 {
	return discount.Apply(price)
}

func main() {
	checkout := Checkout{}

	fmt.Println("Precio sin descuento:", checkout.FinalPrice(100, NoDiscount{}))
	fmt.Println("Precio con 10% de descuento:", checkout.FinalPrice(100, TenPercentDiscount{}))
}

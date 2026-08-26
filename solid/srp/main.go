// SRP - Principio de responsabilidad única
package main

import "fmt"

// Product representa un producto del carrito.
type Product struct {
	Name  string
	Price float64
}

// ShoppingCart tiene una responsabilidad: mantener los productos y calcular
// el total del carrito.
type ShoppingCart struct {
	products []Product
}

func (c *ShoppingCart) Add(product Product) {
	c.products = append(c.products, product)
}

func (c ShoppingCart) Total() float64 {
	total := 0.0
	for _, product := range c.products {
		total += product.Price
	}
	return total
}

// ReceiptPrinter tiene otra responsabilidad: mostrar el total de una compra.
// No necesita conocer cómo se guardan los productos ni cómo se calcula el total.
type ReceiptPrinter struct{}

func (ReceiptPrinter) Print(cart ShoppingCart) {
	fmt.Println("Total de la compra:", cart.Total())
}

func main() {
	cart := ShoppingCart{}
	cart.Add(Product{Name: "Libro", Price: 20})
	cart.Add(Product{Name: "Cuaderno", Price: 5})

	printer := ReceiptPrinter{}
	printer.Print(cart)
}

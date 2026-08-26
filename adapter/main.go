// Adapter permite usar una interfaz existente desde una interfaz esperada.
package main

import "fmt"

type PaymentProcessor interface {
	Pay(amount float64)
}

// LegacyBank tiene una API que no coincide con la que necesita la aplicación.
type LegacyBank struct{}

func (LegacyBank) MakeTransfer(amount float64) {
	fmt.Println("Transferencia bancaria:", amount)
}

// BankAdapter traduce PaymentProcessor a la API de LegacyBank.
type BankAdapter struct {
	bank LegacyBank
}

func (a BankAdapter) Pay(amount float64) {
	a.bank.MakeTransfer(amount)
}

func Checkout(processor PaymentProcessor, amount float64) {
	processor.Pay(amount)
}

func main() {
	Checkout(BankAdapter{bank: LegacyBank{}}, 150)
}

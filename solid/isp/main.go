// ISP - Principio de segregación de interfaces
package main

import "fmt"

// Una interfaz describe una capacidad pequeña y concreta.
type Printer interface {
	Print(document string)
}

type Scanner interface {
	Scan(document string)
}

// BasicPrinter solo imprime. No tiene que implementar Scan porque no depende
// de una interfaz grande que mezcle capacidades diferentes.
type BasicPrinter struct{}

func (BasicPrinter) Print(document string) {
	fmt.Println("Imprimiendo:", document)
}

// OfficeDevice tiene las dos capacidades y, por eso, implementa ambas
// interfaces de manera implícita.
type OfficeDevice struct{}

func (OfficeDevice) Print(document string) {
	fmt.Println("Imprimiendo:", document)
}

func (OfficeDevice) Scan(document string) {
	fmt.Println("Escaneando:", document)
}

// Esta función solo necesita imprimir y no necesita saber si recibe una
// BasicPrinter o un OfficeDevice.
func PrintDocument(printer Printer, document string) {
	printer.Print(document)
}

// Esta interfaz compuesta sirve para un cliente que necesita ambas capacidades.
// Un tipo que la implementa debe tener Print y Scan.
type MultiFunctionDevice interface {
	Printer
	Scanner
}

func CopyDocument(device MultiFunctionDevice, document string) {
	device.Scan(document)
	device.Print(document)
}

func main() {
	document := "informe.txt"

	PrintDocument(BasicPrinter{}, document)

	officeDevice := OfficeDevice{}
	PrintDocument(officeDevice, document)
	CopyDocument(officeDevice, document)
}

// ISP - Principio de segregación de interfaces
package main

import "fmt"

// El ISP establece que ningún cliente debería verse obligado a depender de
// métodos que no utiliza. En este ejemplo, la interfaz Printer tiene tres
// métodos: Print, Fax y Scan. Sin embargo, no todas las impresoras tienen
// capacidades de fax o escaneo. Por lo tanto, podemos crear interfaces
// separadas para cada capacidad.

type Document struct {
}

type Printer interface {
	Print(d Document) error
	Fax(d Document) error
	Scan(d Document) error
}

type MultiFunctionPrinter struct {
}

func (mfp MultiFunctionPrinter) Print(d Document) error {
	fmt.Println("Printing document...")
	return nil
}

func (mfp MultiFunctionPrinter) Fax(d Document) error {
	fmt.Println("Faxing document...")
	return nil
}

func (mfp MultiFunctionPrinter) Scan(d Document) error {
	fmt.Println("Scanning document...")
	return nil
}

type OldFashionedPrinter struct {
}

func (ofp OldFashionedPrinter) Print(d Document) error {
	fmt.Println("Printing document...")
	return nil
}

// OldFashionedPrinter no tiene capacidades de fax ni de escaneo,
// por lo que no debería verse obligado a implementar esos métodos.

// ISP
// Ahora podemos crear interfaces separadas para cada capacidad.
type PrinterOnly interface {
	Print(d Document) error
}

type FaxOnly interface {
	Fax(d Document) error
}

type ScanOnly interface {
	Scan(d Document) error
}

func main() {
	doc := Document{}
	var p Printer = MultiFunctionPrinter{}
	var po PrinterOnly = MultiFunctionPrinter{}

	p.Print(doc)
	p.Scan(doc)
	p.Fax(doc)

	po.Print(doc)

}

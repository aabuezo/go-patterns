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

type PrinterScanner interface {
	Print(d Document) error
	Scan(d Document) error
}

type NewerPrinter struct {
}

func (mfp NewerPrinter) Print(d Document) error {
	fmt.Println("Printing document...")
	return nil
}

func (mfp NewerPrinter) Scan(d Document) error {
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
// por lo que no debería verse obligado a implementar Scan.

// ISP
// Ahora podemos crear interfaces separadas para cada capacidad.
type Printer interface {
	Print(d Document) error
}

type Scanner interface {
	Scan(d Document) error
}

// Tambien podemos crear una interfaz que incluya ambas capacidades,
// para aquellos clientes que necesiten ambas.
type MultiFunctionDevice interface {
	Printer
	Scanner
}

func main() {
	doc := Document{}

	mfd := NewerPrinter{}
	mfd.Print(doc)
	mfd.Scan(doc)

	p := OldFashionedPrinter{}
	p.Print(doc)

}

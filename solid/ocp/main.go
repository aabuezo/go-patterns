// OCP - Principio abierto/cerrado
package main

import "fmt"

// OCP: abierto para la extensión, cerrado para la modificación.
// Specification es un patrón empresarial.
// Esto significa que, una vez que diseñaste y probaste la API de un tipo,
// no deberías tener que modificarla, pero sí puedes extenderla agregando
// nuevas funcionalidades. Ese tipo ya funciona y tienes clientes que
// dependen de él; no quieres romperlos cambiando la API.
type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct{}

// Los siguientes métodos no están abiertos a la extensión ni cerrados a la
// modificación. Violan el principio OCP porque, si queremos agregar un
// filtro nuevo, tenemos que modificar la estructura Filter.
func (f Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// Los siguientes métodos no están abiertos a la extensión ni cerrados a la modificación.
func (f Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

// Los siguientes métodos no están abiertos a la extensión ni cerrados a la modificación.
func (f Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// El patrón Specification es una mejor forma de implementar el principio OCP:
// está abierto a la extensión y cerrado a la modificación.
type Specification interface {
	// Indica si un producto cumple un criterio determinado.
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

// Especificación compuesta que combina dos especificaciones.
type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

// Si quieres filtrar por un criterio nuevo, puedes crear una especificación
// que implemente la interfaz Specification sin modificar el código existente.

// Es poco probable que alguna vez modifiques esta estructura,
// pero puedes extenderla agregando nuevas especificaciones.
type BetterFilter struct{}

// Un único método que recibe una especificación y filtra los productos según ella.
func (b BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	f := Filter{}

	fmt.Println("Green products (old):")
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	// Usando el nuevo filtro con especificaciones.
	bf := BetterFilter{}
	greenSpec := ColorSpecification{green}
	largeSpec := SizeSpecification{large}

	fmt.Println("Green products (new):")
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Println("Large products (new):")
	for _, v := range bf.Filter(products, largeSpec) {
		fmt.Printf(" - %s is large\n", v.name)
	}

	// Combinando especificaciones.
	greenAndLargeSpec := AndSpecification{greenSpec, largeSpec}
	fmt.Println("Green and large products (new):")
	for _, v := range bf.Filter(products, greenAndLargeSpec) {
		fmt.Printf(" - %s is green and large\n", v.name)
	}
}

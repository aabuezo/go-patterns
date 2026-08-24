// OCP - Open Responsibility Principle
package main

import "fmt"

// OCP: Open for extension, closed for modification
// Specification: is an enterprise pattern
// it means that once you designed and tested the API of a particular type,
// you shouldn't really have to modify it, but you can extend it by adding new functionality
// you already got that type working, you already have clients relying on it,
// you don't want to break those clients by changing the API
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

// the following methods are not open for extension, closed for modification
// they violate the OCP principle because if we want to add a new filter,
// we have to modify the Filter struct
func (f Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// the following methods are not open for extension, closed for modification
func (f Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

// the following methods are not open for extension, closed for modification
func (f Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// The Specification pattern is a better way to implement the OCP principle
// it is open for extension, closed for modification
type Specification interface {
	// specifies whether a product satisfies a certain criteria
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

// composite specification that combines two specifications
type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

// if you want to filter by a new criteria, you can create a new specification
// that implements the Specification interface, without modifying the existing code

// you're unlikely to ever modify this struct,
// but you can extend it by adding new specifications
type BetterFilter struct{}

// Just one method that takes a specification and filters the products based on it
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

	// using the new filter with specifications
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

	// combining specifications
	greenAndLargeSpec := AndSpecification{greenSpec, largeSpec}
	fmt.Println("Green and large products (new):")
	for _, v := range bf.Filter(products, greenAndLargeSpec) {
		fmt.Printf(" - %s is green and large\n", v.name)
	}
}

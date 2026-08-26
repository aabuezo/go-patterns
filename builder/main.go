// Builder construye un objeto paso a paso cuando tiene varias opciones.
package main

import "fmt"

type Sandwich struct {
	bread   string
	cheese  bool
	tomato  bool
	lettuce bool
}

type SandwichBuilder struct {
	sandwich Sandwich
}

func (b *SandwichBuilder) WithBread(bread string) *SandwichBuilder {
	b.sandwich.bread = bread
	return b
}

func (b *SandwichBuilder) AddCheese() *SandwichBuilder {
	b.sandwich.cheese = true
	return b
}

func (b *SandwichBuilder) AddTomato() *SandwichBuilder {
	b.sandwich.tomato = true
	return b
}

func (b *SandwichBuilder) AddLettuce() *SandwichBuilder {
	b.sandwich.lettuce = true
	return b
}

func (b *SandwichBuilder) Build() Sandwich {
	return b.sandwich
}

func main() {
	sandwich := new(SandwichBuilder).
		WithBread("integral").
		AddCheese().
		AddTomato().
		Build()

	fmt.Printf("Sandwich: %+v\n", sandwich)
}

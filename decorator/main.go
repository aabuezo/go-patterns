// Decorator agrega comportamiento a un objeto sin modificar su tipo original.
package main

import "fmt"

type Coffee interface {
	Description() string
	Cost() float64
}

type SimpleCoffee struct{}

func (SimpleCoffee) Description() string { return "café" }
func (SimpleCoffee) Cost() float64       { return 2 }

type MilkDecorator struct {
	coffee Coffee
}

func (d MilkDecorator) Description() string {
	return d.coffee.Description() + " con leche"
}

func (d MilkDecorator) Cost() float64 {
	return d.coffee.Cost() + 0.5
}

func main() {
	coffee := MilkDecorator{coffee: SimpleCoffee{}}
	fmt.Println(coffee.Description(), coffee.Cost())
}

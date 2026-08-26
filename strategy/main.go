// Strategy permite elegir un algoritmo sin cambiar el código que lo utiliza.
package main

import "fmt"

type ShippingStrategy interface {
	Cost(distance int) float64
}

type StandardShipping struct{}

func (StandardShipping) Cost(distance int) float64 {
	return float64(distance) * 0.5
}

type ExpressShipping struct{}

func (ExpressShipping) Cost(distance int) float64 {
	return float64(distance) * 1.5
}

type Delivery struct {
	strategy ShippingStrategy
}

func (d Delivery) Cost(distance int) float64 {
	return d.strategy.Cost(distance)
}

func main() {
	standard := Delivery{strategy: StandardShipping{}}
	express := Delivery{strategy: ExpressShipping{}}

	fmt.Println("Envío estándar:", standard.Cost(10))
	fmt.Println("Envío express:", express.Cost(10))
}

// LSP - Principio de sustitución de Liskov
package main

import "fmt"

// Bird representa lo que todos los pájaros de este ejemplo pueden hacer.
// No se incluye Fly aquí porque no todos los pájaros pueden volar.
type Bird interface {
	Eat()
}

type FlyingBird interface {
	Bird
	Fly()
}

type Sparrow struct{}

func (Sparrow) Eat() {
	fmt.Println("El gorrión come")
}

func (Sparrow) Fly() {
	fmt.Println("El gorrión vuela")
}

type Penguin struct{}

func (Penguin) Eat() {
	fmt.Println("El pingüino come")
}

// Esta función solo recibe pájaros que cumplen el contrato de volar.
func MakeBirdFly(bird FlyingBird) {
	bird.Fly()
}

func main() {
	var bird Bird = Penguin{}
	bird.Eat()

	// Sparrow cumple Bird y FlyingBird.
	var flyingBird FlyingBird = Sparrow{}
	flyingBird.Eat()
	MakeBirdFly(flyingBird)
}

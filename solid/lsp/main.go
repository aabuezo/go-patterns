// LSP - Principio de sustitución de Liskov
package main

import "fmt"

// LSP: los subtipos deben poder sustituir a sus tipos base.
// Esto significa que, si S es un subtipo de T, los objetos de tipo T pueden
// reemplazarse por objetos de tipo S sin alterar las propiedades deseables
// del programa (corrección, tarea realizada, etc.).
type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

// Aquí, Square es un subtipo de Rectangle, pero viola el principio LSP porque
// cambia el comportamiento del método Area. El método Area de Square no
// devuelve el área de un rectángulo, sino el área de un cuadrado. Esto puede
// producir comportamientos inesperados al usar Square en lugar de Rectangle.
type Square struct {
	Rectangle
}

func (s *Square) SetWidth(width float64) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height float64) {
	s.width = height
	s.height = height
}

// La forma correcta de implementar el principio LSP es crear una nueva
// interfaz para las figuras que tienen un área y luego implementarla tanto
// para Rectangle como para Square. De esta manera, podemos usar Rectangle y
// Square indistintamente sin violar el principio LSP.

type Shape interface {
	Area() float64
}

func (s *Square) Area() float64 {
	return s.width * s.height
}

func main() {
	r := Rectangle{width: 5, height: 10}
	fmt.Println("Rectangle Area:", r.Area())

	s := Square{}
	s.SetWidth(5)
	fmt.Println("Square Area:", s.Area())
}

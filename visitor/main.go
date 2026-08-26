// Visitor agrega operaciones a una estructura sin modificar sus tipos.
package main

import "fmt"

type Shape interface {
	Accept(visitor ShapeVisitor)
}

type ShapeVisitor interface {
	VisitCircle(circle Circle)
	VisitRectangle(rectangle Rectangle)
}

type Circle struct {
	radius float64
}

func (c Circle) Accept(visitor ShapeVisitor) {
	visitor.VisitCircle(c)
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Accept(visitor ShapeVisitor) {
	visitor.VisitRectangle(r)
}

// AreaVisitor agrega la operación de calcular áreas sin modificar Circle
// ni Rectangle.
type AreaVisitor struct{}

func (AreaVisitor) VisitCircle(circle Circle) {
	fmt.Println("Área del círculo:", 3.14*circle.radius*circle.radius)
}

func (AreaVisitor) VisitRectangle(rectangle Rectangle) {
	fmt.Println("Área del rectángulo:", rectangle.width*rectangle.height)
}

func main() {
	visitor := AreaVisitor{}
	shapes := []Shape{
		Circle{radius: 2},
		Rectangle{width: 3, height: 4},
	}

	for _, shape := range shapes {
		shape.Accept(visitor)
	}
}

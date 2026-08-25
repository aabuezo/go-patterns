// LSP - Liskov Substitution Principle
package main

import "fmt"

// LSP: Subtypes must be substitutable for their base types
// it means that if S is a subtype of T, then objects of type T may be replaced with objects of type S
// without altering any of the desirable properties of that program (correctness, task performed, etc.)
type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

// Here, Square is a subtype of Rectangle, but it violates the LSP principle because it changes the
// behavior of the Area method. The Area method of Square does not return the area of a rectangle,
// but rather the area of a square. This can lead to unexpected behavior when using Square in place of Rectangle.
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

// The correct way to implement the LSP principle is to create a new interface for shapes that have an area,
// and then implement that interface for both Rectangle and Square. This way, we can use both Rectangle and
// Square interchangeably without violating the LSP principle.

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

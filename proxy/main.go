// Proxy controla el acceso a otro objeto.
package main

import "fmt"

type Image interface {
	Display()
}

type RealImage struct {
	file string
}

func (i RealImage) Display() {
	fmt.Println("Mostrando imagen:", i.file)
}

type ImageProxy struct {
	file  string
	image *RealImage
}

func (p *ImageProxy) Display() {
	if p.image == nil {
		fmt.Println("Cargando imagen:", p.file)
		p.image = &RealImage{file: p.file}
	}
	p.image.Display()
}

func main() {
	image := &ImageProxy{file: "foto.jpg"}
	image.Display()
	image.Display()
}

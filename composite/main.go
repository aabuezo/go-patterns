// Composite permite tratar un objeto individual y un grupo de objetos igual.
package main

import "fmt"

type FileSystemItem interface {
	Size() int
}

type File struct {
	name string
	size int
}

func (f File) Size() int { return f.size }

type Folder struct {
	name  string
	items []FileSystemItem
}

func (f Folder) Size() int {
	total := 0
	for _, item := range f.items {
		total += item.Size()
	}
	return total
}

func main() {
	folder := Folder{
		name: "proyecto",
		items: []FileSystemItem{
			File{name: "main.go", size: 10},
			File{name: "README.md", size: 5},
		},
	}

	fmt.Println("Tamaño de la carpeta:", folder.Size())
}

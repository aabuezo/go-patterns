// SRP = Principio de responsabilidad única
package main

// Un tipo debería tener una única responsabilidad.

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

var entryCount = 0

// El tipo Journal tiene una única responsabilidad (separación de intereses):
// agregar, eliminar, etc. entradas y administrar esas entradas.
type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s",
		entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// Separación de intereses.
// ¡La persistencia no es responsabilidad de Journal!
// La persistencia puede ser administrada por otro componente.
// Pensemos también en otros tipos que necesitan persistencia.
// ¡La persistencia es común a muchos objetos!
func (j *Journal) Save(filename string) { // ¡Rompe el SRP!
	_ = os.WriteFile(filename,
		[]byte(j.String()), os.ModeAppend)
}

func (j *Journal) Load(filename string) { // ¡Rompe el SRP!

}

func (j *Journal) LoadFromWeb(url *url.URL) { // ¡Rompe el SRP!

}

// Separar los intereses podría verse de la siguiente manera
// (a nivel de paquete).
var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename,
		[]byte(strings.Join(j.entries, LineSeparator)), 0644)
}

// O de esta manera (usando un objeto separado).
type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate a bug")
	fmt.Println(j.String())

	// Persistencia.
	SaveToFile(&j, "journal.txt")
	//
	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")
}

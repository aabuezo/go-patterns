// SRP = Single Responsibiility Principle
package main

// A type should have one responsibility only.

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

var entryCount = 0

// type Journal has a single responsibility (separation of concerns):
// add/remove/etc of keeping entries and managing those entries
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

// separation of concerns
// the responsibility of the Journal is not persistance!
// Persistance can be managed by a different component
// Think about other types that need persistance as well
// persistance is common to many objects!
func (j *Journal) Save(filename string) { // Breaks SRP!!
	_ = os.WriteFile(filename,
		[]byte(j.String()), os.ModeAppend)
}

func (j *Journal) Load(filename string) { // Breaks SRP!!

}

func (j *Journal) LoadFromWeb(url *url.URL) { // Breaks SRP!!

}

// Separating concerns would look like something like this
// (package level)
var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename,
		[]byte(strings.Join(j.entries, LineSeparator)), 0644)
}

// or this (using a separate object)
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

	// Persistence
	SaveToFile(&j, "journal.txt")
	//
	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")
}

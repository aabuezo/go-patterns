// State permite que un objeto cambie su comportamiento según su estado actual.
package main

import "fmt"

type DoorState interface {
	Open(door *Door)
	Close(door *Door)
}

type Door struct {
	state DoorState
}

func (d *Door) Open()  { d.state.Open(d) }
func (d *Door) Close() { d.state.Close(d) }

type ClosedDoor struct{}

func (ClosedDoor) Open(door *Door) {
	fmt.Println("Puerta abierta")
	door.state = OpenDoor{}
}

func (ClosedDoor) Close(*Door) {
	fmt.Println("La puerta ya está cerrada")
}

type OpenDoor struct{}

func (OpenDoor) Open(*Door) {
	fmt.Println("La puerta ya está abierta")
}

func (OpenDoor) Close(door *Door) {
	fmt.Println("Puerta cerrada")
	door.state = ClosedDoor{}
}

func main() {
	door := Door{state: ClosedDoor{}}
	door.Close()
	door.Open()
	door.Open()
	door.Close()
}

// Command convierte una acción en un objeto que puede ejecutarse o guardarse.
package main

import "fmt"

type Command interface {
	Execute()
}

type Light struct{}

func (Light) TurnOn() {
	fmt.Println("Luz encendida")
}

type TurnOnLightCommand struct {
	light Light
}

func (c TurnOnLightCommand) Execute() {
	c.light.TurnOn()
}

type RemoteControl struct{}

func (RemoteControl) Press(command Command) {
	command.Execute()
}

func main() {
	remote := RemoteControl{}
	remote.Press(TurnOnLightCommand{light: Light{}})
}

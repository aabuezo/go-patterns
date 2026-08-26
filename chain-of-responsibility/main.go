// Chain of Responsibility pasa una solicitud por una cadena de handlers.
package main

import "fmt"

type SupportRequest struct {
	message string
	urgent  bool
}

type Handler interface {
	Handle(request SupportRequest) bool
}

type FirstLevelSupport struct{}

func (FirstLevelSupport) Handle(request SupportRequest) bool {
	if !request.urgent {
		fmt.Println("Soporte de primer nivel:", request.message)
		return true
	}
	return false
}

type ManagerSupport struct{}

func (ManagerSupport) Handle(request SupportRequest) bool {
	fmt.Println("Manager atiende:", request.message)
	return true
}

func ProcessRequest(request SupportRequest, handlers ...Handler) {
	for _, handler := range handlers {
		if handler.Handle(request) {
			return
		}
	}
}

func main() {
	ProcessRequest(SupportRequest{"No puedo iniciar sesión", false}, FirstLevelSupport{}, ManagerSupport{})
	ProcessRequest(SupportRequest{"El sistema está caído", true}, FirstLevelSupport{}, ManagerSupport{})
}

// Singleton garantiza una única instancia compartida de un componente.
package main

import (
	"fmt"
	"sync"
)

type AppConfig struct {
	Environment string
}

var (
	config     *AppConfig
	configOnce sync.Once
)

// GetConfig crea la configuración una sola vez y luego devuelve la misma instancia.
func GetConfig() *AppConfig {
	configOnce.Do(func() {
		config = &AppConfig{Environment: "development"}
	})
	return config
}

func main() {
	first := GetConfig()
	second := GetConfig()

	fmt.Println(first.Environment)
	fmt.Println(first == second)
}

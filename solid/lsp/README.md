# LSP en Go: principio de sustitución de Liskov

LSP significa **Liskov Substitution Principle**.

La idea central es:

> Si un tipo se presenta como una variante de otro tipo, cualquier código que espera el tipo original debería seguir funcionando correctamente al recibir la variante.

No alcanza con que la variante tenga métodos parecidos. También debe respetar el comportamiento que el código cliente espera.

## El ejemplo del rectángulo y el cuadrado

Un rectángulo permite cambiar su ancho y su alto independientemente. Por ejemplo:

```go
r.SetWidth(5)
r.SetHeight(10)
```

Después de esas operaciones esperamos que el área sea `50`.

Un cuadrado, en cambio, necesita que ancho y alto sean siempre iguales. Si recibe las mismas operaciones:

```go
s.SetWidth(5)  // ancho = 5, alto = 5
s.SetHeight(10) // ancho = 10, alto = 10
```

el área termina siendo `100`, no `50`.

Si `Square` pudiera sustituir a `Rectangle`, el código que trabaja con un rectángulo recibiría un resultado inesperado. Esa es la violación de LSP: la variante no respeta las expectativas del tipo original.

## Importante en Go

Go no tiene herencia clásica. Esto:

```go
type Square struct {
    Rectangle
}
```

es **embedding** o composición. `Square` contiene un `Rectangle` y puede recibir algunos métodos promocionados, pero no se convierte automáticamente en un subtipo de `Rectangle`.

Por ejemplo, una función que recibe un `Rectangle` no acepta directamente un `Square`:

```go
func printArea(r Rectangle) {
    fmt.Println(r.Area())
}
```

La violación de LSP en este ejercicio es principalmente conceptual: estamos modelando al cuadrado como una clase especial de rectángulo, aunque sus reglas para modificar ancho y alto son incompatibles.

## La solución con interfaces

En tu código, esta interfaz representa una propiedad común válida:

```go
type Shape interface {
    Area() float64
}
```

Una función que solo necesita calcular áreas puede trabajar con ambos tipos:

```go
func printArea(s Shape) {
    fmt.Println(s.Area())
}
```

Tanto `Rectangle` como `Square` cumplen el contrato de `Shape`, y ninguno rompe las expectativas de la interfaz. La interfaz solo promete que el objeto sabe calcular su área; no promete que se pueda modificar el ancho y el alto de una manera determinada.

## Regla práctica

Cuando diseñes una interfaz o una relación entre tipos, preguntate:

> ¿Cualquier implementación puede cumplir este contrato sin producir sorpresas para quien la usa?

Si la respuesta es sí, probablemente estás respetando LSP.

En este ejemplo, `Rectangle` y `Square` pueden compartir `Shape`, pero no deberían compartir una abstracción que permita cambiar ancho y alto libremente.


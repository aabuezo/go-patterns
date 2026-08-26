# Single Responsibility Principle (SRP)

El principio de responsabilidad única dice:

> Un tipo debería tener una sola responsabilidad y, por lo tanto, una sola razón importante para cambiar.

Una responsabilidad es un grupo de tareas relacionadas. No significa necesariamente que el tipo tenga un solo método.

## El ejemplo del carrito

`ShoppingCart` se ocupa de mantener los productos y calcular su total:

```go
type ShoppingCart struct {
	products []Product
}
```

Sus métodos `Add` y `Total` pertenecen a la misma responsabilidad: administrar el contenido del carrito.

`ReceiptPrinter` tiene otra responsabilidad: mostrar el total al usuario.

```go
type ReceiptPrinter struct{}

func (ReceiptPrinter) Print(cart ShoppingCart) {
	fmt.Println("Total de la compra:", cart.Total())
}
```

## ¿Qué problema evitamos?

Si `ShoppingCart` también imprimiera recibos, guardara archivos y enviara emails, tendría varias razones para cambiar:

- cambiaría si cambia la forma de calcular el total;
- cambiaría si cambia el formato del recibo;
- cambiaría si cambia la forma de enviar el recibo.

Separando esas tareas, cada tipo puede modificarse por un motivo claro sin afectar responsabilidades que no le corresponden.

## Regla práctica

Preguntate:

> ¿Cuántos motivos diferentes existen para modificar este tipo?

Si la respuesta incluye tareas de áreas diferentes, probablemente el tipo tenga más de una responsabilidad.

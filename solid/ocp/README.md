# Open/Closed Principle (OCP)

El principio abierto/cerrado dice:

> Un componente debería estar abierto para extenderse, pero cerrado para modificarse.

Esto significa que, cuando aparece un comportamiento nuevo, idealmente podemos agregar una nueva implementación sin cambiar el código que ya funciona.

## El ejemplo de los descuentos

`Checkout` necesita aplicar un descuento, pero no necesita conocer todos los tipos posibles:

```go
type Discount interface {
	Apply(price float64) float64
}
```

El método `FinalPrice` recibe cualquier `Discount`:

```go
func (Checkout) FinalPrice(price float64, discount Discount) float64 {
	return discount.Apply(price)
}
```

Para agregar un descuento nuevo, por ejemplo del 20%, creamos un tipo que implemente `Discount`:

```go
type TwentyPercentDiscount struct{}

func (TwentyPercentDiscount) Apply(price float64) float64 {
	return price * 0.80
}
```

No necesitamos modificar `Checkout`, `FinalPrice`, `NoDiscount` ni `TenPercentDiscount`.

## El error que queremos evitar

Un diseño menos flexible tendría un método con un `switch`:

```go
func FinalPrice(price float64, discountType string) float64 {
	switch discountType {
	case "10%":
		return price * 0.90
	case "20%":
		return price * 0.80
	}
	return price
}
```

Cada nuevo descuento obligaría a modificar este método. Con una interfaz, el nuevo comportamiento se agrega como una extensión.

## Regla práctica

Si cada nuevo caso te obliga a agregar otro `if`, `switch` o método al componente principal, preguntate si ese comportamiento variable podría representarse mediante una interfaz.

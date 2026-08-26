# Open/Closed Principle (OCP)

El principio abierto/cerrado dice:

> Un componente debería estar abierto para extenderse, pero cerrado para modificarse.

Esto significa que, cuando aparece un comportamiento nuevo, idealmente podemos agregar una nueva implementación sin cambiar el código que ya funciona.

En este ejemplo, el componente que queremos mantener cerrado es `Checkout`. `main` cumple otro rol: conecta las piezas y decide qué descuento usar. Por eso es normal modificar `main` para probar o seleccionar un descuento nuevo. Lo importante es que no tenemos que modificar `Checkout` ni `FinalPrice` cada vez que aparece una nueva regla.

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

La ventaja se nota más cuando el descuento no consiste solamente en multiplicar por un porcentaje. Por ejemplo, Black Friday podría aplicar una tasa distinta según el precio:

```go
type BlackFridayDiscount struct{}

func (BlackFridayDiscount) Apply(price float64) float64 {
	if price > 1000 {
		return price * 0.70
	}
	return price * 0.85
}
```

Para usarlo, `main` agrega la selección de esta nueva estrategia:

```go
fmt.Println("Precio de Black Friday:", checkout.FinalPrice(1500, BlackFridayDiscount{}))
```

No fue necesario agregar otro `if` a `FinalPrice`. `Checkout` solo sabe que recibió algo con un método `Apply`; no necesita conocer las reglas de Black Friday, del descuento del 10% ni de ningún descuento futuro.

Si todos los descuentos fueran siempre porcentajes fijos, pasar un `float64` podría ser más simple. La interfaz resulta útil cuando lo que cambia no es solamente un valor, sino la lógica para calcular el precio.

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

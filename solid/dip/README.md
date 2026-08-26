# Dependency Inversion Principle (DIP)

El principio de inversión de dependencias establece que:

- Los módulos de alto nivel no deberían depender directamente de módulos de bajo nivel.
- Ambos deberían depender de abstracciones, normalmente interfaces.

Un módulo de alto nivel contiene reglas importantes del negocio. Un módulo de bajo nivel se ocupa de detalles concretos, como enviar un email, guardar datos o escribir en un archivo.

En el ejemplo, `OrderService` es el módulo de alto nivel porque confirma una orden. `EmailSender` es un módulo de bajo nivel porque conoce los detalles de enviar un email.

`OrderService` no depende directamente de `EmailSender`. Depende de la interfaz `MessageSender`:

```go
type MessageSender interface {
	Send(message string)
}
```

De esta forma, cualquier tipo que implemente `MessageSender` puede utilizarse como dependencia. Por ejemplo, en el futuro podríamos crear un `SMSSender` sin modificar `OrderService`.

La función `NewOrderService` recibe el objeto que debe utilizar. Esto se llama inyección de dependencias: la dependencia se entrega desde afuera en lugar de crearla dentro de `OrderService`.

El beneficio es que el módulo de alto nivel queda menos acoplado a una implementación concreta y resulta más fácil de cambiar y probar.

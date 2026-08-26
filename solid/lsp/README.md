# Liskov Substitution Principle (LSP)

El principio de sustitución de Liskov dice:

> Si una función espera un tipo o una interfaz, cualquier implementación válida debería poder reemplazarlo sin romper el comportamiento esperado.

Una forma simple de detectarlo es revisar qué promete una abstracción. No deberíamos incluir en un contrato una capacidad que algunas implementaciones no pueden cumplir.

## El ejemplo de los pájaros

Todos los pájaros del ejemplo pueden comer, por eso `Bird` solo declara `Eat`:

```go
type Bird interface {
	Eat()
}
```

Un gorrión también puede volar, así que puede cumplir una interfaz más específica:

```go
type FlyingBird interface {
	Bird
	Fly()
}
```

El pingüino cumple `Bird`, pero no cumple `FlyingBird`. Por eso no se lo puede pasar a `MakeBirdFly`, que necesita un pájaro capaz de volar.

## El error que queremos evitar

Una interfaz como esta sería problemática:

```go
type Bird interface {
	Eat()
	Fly()
}
```

Obligaría al pingüino a implementar `Fly`, aunque no puede hacerlo. Una implementación que devuelve un error, no hace nada o produce un comportamiento inesperado no estaría respetando el contrato.

## Regla práctica

Preguntate:

> ¿Puedo reemplazar una implementación por otra sin que el código que la usa se sorprenda?

Si no, el contrato probablemente es demasiado amplio o la abstracción no representa correctamente a todas sus implementaciones.

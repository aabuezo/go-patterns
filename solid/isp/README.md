# Interface Segregation Principle (ISP)

El principio de segregación de interfaces dice:

> Ningún cliente debería depender de métodos que no necesita.

La idea es crear interfaces pequeñas, donde cada una represente una capacidad concreta.

## El ejemplo

En `main.go` separamos las capacidades de una máquina:

```go
type Printer interface {
	Print(document string)
}

type Scanner interface {
	Scan(document string)
}
```

`BasicPrinter` solo imprime, así que implementa `Printer` y no necesita tener un método `Scan` inútil.

`OfficeDevice` puede imprimir y escanear, por lo que implementa ambas interfaces. En Go, esto ocurre de manera implícita: alcanza con que el tipo tenga los métodos requeridos.

La función `PrintDocument` recibe un `Printer` porque solo necesita imprimir. Puede trabajar tanto con `BasicPrinter` como con `OfficeDevice` sin conocer sus detalles.

## Interfaces compuestas

Si una función necesita las dos capacidades, podemos combinar las interfaces:

```go
type MultiFunctionDevice interface {
	Printer
	Scanner
}
```

`CopyDocument` usa esta interfaz porque necesita escanear e imprimir.

## Regla práctica

Antes de crear una interfaz grande, preguntate:

> ¿Todos los tipos que la implementen realmente necesitan todos estos métodos?

Si la respuesta es no, probablemente convenga dividirla en interfaces más pequeñas.

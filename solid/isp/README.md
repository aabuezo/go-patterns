# Interface Segregation Principle (ISP)

El principio de segregación de interfaces establece que ningún cliente debería depender de métodos que no necesita.

En lugar de crear una interfaz grande con muchas responsabilidades, conviene dividirla en interfaces pequeñas y específicas. Así, cada tipo o función puede depender únicamente de las operaciones que realmente utiliza.

Por ejemplo, una impresora que solo imprime no debería verse obligada a implementar el método `Scan`. Podemos separar cada capacidad en una interfaz pequeña:

```go
type Printer interface {
	Print(d Document) error
}

type Scanner interface {
	Scan(d Document) error
}
```

`OldFashionedPrinter` implementa únicamente `Printer`, mientras que `NewerPrinter` implementa tanto `Printer` como `Scanner`. En Go, la implementación de interfaces es implícita: un tipo cumple una interfaz cuando posee todos sus métodos.

También podemos crear una interfaz compuesta que incluya otras interfaces:

```go
type MultiFunctionDevice interface {
	Printer
	Scanner
}
```

Todo tipo que implemente `MultiFunctionDevice` debe implementar los métodos de `Printer` y `Scanner`, es decir, `Print` y `Scan`.

El objetivo del ISP es reducir el acoplamiento y hacer que el código sea más simple de implementar, mantener y probar.

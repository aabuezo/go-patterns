# Interface Segregation Principle (ISP)

El principio de segregación de interfaces establece que ningún cliente debería depender de métodos que no necesita.

En lugar de crear una interfaz grande con muchas responsabilidades, conviene dividirla en interfaces pequeñas y específicas. Así, cada tipo o función puede depender únicamente de las operaciones que realmente utiliza.

Por ejemplo, una impresora que solo imprime no debería verse obligada a implementar métodos como `Scan` o `Fax`. Podemos definir una interfaz pequeña:

```go
type PrinterOnly interface {
	Print()
}
```

Una impresora multifunción puede implementar también otras interfaces, porque tiene más capacidades. En Go, la implementación de interfaces es implícita: un tipo cumple una interfaz cuando posee todos sus métodos.

El objetivo del ISP es reducir el acoplamiento y hacer que el código sea más simple de implementar, mantener y probar.

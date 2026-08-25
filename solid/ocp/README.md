# OCP en Go: principio abierto/cerrado

OCP significa **Open/Closed Principle**.

La idea central es:

> Un componente debería estar abierto para extenderse, pero cerrado para modificarse.

Esto significa que, cuando aparece un nuevo comportamiento, idealmente podemos agregar código nuevo sin tener que cambiar y arriesgar el código existente que ya funciona.

## El problema en `Filter`

En tu ejemplo, `Filter` tiene métodos separados:

```go
FilterByColor(...)
FilterBySize(...)
FilterBySizeAndColor(...)
```

Mientras solo existan esos criterios, funciona. Pero si necesitás filtrar por precio, nombre o una combinación nueva, tenés que agregar otro método o modificar la estructura de `Filter`.

Con muchos criterios, `Filter` empieza a acumular casos especiales. El componente principal deja de estar cerrado para modificación: cada filtro nuevo exige tocar código existente.

## La extensión mediante `Specification`

La interfaz de tu ejemplo abstrae el criterio de búsqueda:

```go
type Specification interface {
    IsSatisfied(p *Product) bool
}
```

`BetterFilter` solo necesita saber si un producto cumple una especificación:

```go
func (b BetterFilter) Filter(products []Product, spec Specification) []*Product
```

El filtro no conoce los detalles de color, tamaño ni futuros criterios. Cada criterio se agrega creando un nuevo tipo que implemente `Specification`.

Por ejemplo, `ColorSpecification` y `SizeSpecification` son extensiones independientes. `AndSpecification` permite combinarlas sin modificar `BetterFilter`.

El flujo queda así:

```text
nuevo criterio
      |
      v
nueva Specification ---> BetterFilter.Filter
                              |
                              v
                         productos filtrados
```

`BetterFilter` está cerrado para modificación porque su algoritmo no necesita cambiar. Está abierto para extensión porque puede recibir cualquier nueva implementación de `Specification`.

## Qué no significa OCP

OCP no significa que nunca puedas modificar un archivo. Los requisitos cambian y a veces hay que corregir código existente.

Significa que conviene diseñar los puntos variables detrás de una abstracción. En este caso, lo que cambia es el criterio de filtrado, por eso ese comportamiento queda representado por `Specification`.

## Regla práctica

Preguntate:

> Si mañana aparece un nuevo caso de comportamiento, ¿tengo que agregar un `if`, un `switch` o un método especial al componente principal?

Si la respuesta es sí y esperás que aparezcan muchos casos, puede ser útil extraer ese comportamiento a una interfaz o a componentes separados.

En tu ejercicio, agregar una nueva especificación debería requerir crear un tipo nuevo, sin modificar `BetterFilter` ni las especificaciones que ya funcionan.


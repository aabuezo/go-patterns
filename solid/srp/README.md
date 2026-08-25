# SRP en Go: principio de responsabilidad única

SRP significa **Single Responsibility Principle**.

La idea central es:

> Un tipo debería tener una sola responsabilidad y, por lo tanto, una sola razón importante para cambiar.

Una responsabilidad no significa necesariamente “un solo método”. Significa un grupo coherente de tareas relacionadas.

## La responsabilidad de `Journal`

En tu ejemplo, `Journal` administra las entradas del diario:

```go
type Journal struct {
    entries []string
}
```

Estas operaciones pertenecen naturalmente al diario:

- agregar una entrada;
- mantener sus entradas;
- convertirlas en texto.

Por eso `AddEntry` y `String` forman parte de una responsabilidad coherente: administrar el contenido del diario.

## El problema de mezclar persistencia

Estos métodos agregan otra responsabilidad a `Journal`:

```go
func (j *Journal) Save(filename string)
func (j *Journal) Load(filename string)
func (j *Journal) LoadFromWeb(url *url.URL)
```

Ahora `Journal` no solo administra entradas. También conoce:

- cómo guardar en un archivo;
- cómo leer desde un archivo;
- cómo obtener datos desde la web.

Son responsabilidades distintas y cada una puede cambiar por motivos diferentes. Por ejemplo, podrías cambiar el formato del archivo o reemplazar el origen web sin cambiar el concepto de diario.

Si `Journal` hace todo, cualquier cambio en persistencia obliga a modificar y volver a probar el tipo que administra las entradas.

## Separar las responsabilidades

En tu código, `SaveToFile` separa la persistencia de `Journal`:

```go
SaveToFile(&j, "journal.txt")
```

También aparece un tipo específico:

```go
type Persistence struct {
    lineSeparator string
}
```

`Persistence` se ocupa de guardar, mientras que `Journal` se ocupa de sus entradas. Cada tipo tiene una responsabilidad más clara.

La relación puede verse así:

```text
Journal ----------------> administra entradas
    |
    v
Persistence -------------> guarda esas entradas
```

Separar responsabilidades no significa que los tipos no puedan colaborar. Significa que cada uno debe ocuparse de una preocupación concreta.

## Otra observación del ejemplo

`entryCount` es una variable global. Eso hace que el contador sea compartido por todos los diarios y que el comportamiento dependa del estado global del programa.

Para aprender SRP, lo importante es notar que el diario y la persistencia son responsabilidades distintas. Más adelante también podrías revisar si el contador debería pertenecer al `Journal`, pero eso es una decisión separada del principio principal del ejercicio.

## Regla práctica

Preguntate:

> ¿Cuántos motivos diferentes existen para modificar este tipo?

Si un tipo puede cambiar porque cambia el modelo de datos, el formato de almacenamiento y además la comunicación web, probablemente contiene varias responsabilidades.

En este ejemplo, `Journal` debería concentrarse en el diario y otra pieza debería encargarse de guardar o cargar sus datos.


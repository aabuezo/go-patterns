# Builder

`Builder` construye un objeto paso a paso, especialmente cuando tiene muchas opciones.

En el ejemplo, un sándwich puede tener distintos ingredientes. `SandwichBuilder` permite elegirlos con llamadas encadenadas y finalmente crear el objeto con `Build`.

Este patrón evita constructores con muchos parámetros y hace más legible la creación de objetos configurables. En Go también se usan mucho las opciones funcionales, pero Builder resulta útil para aprender la misma idea.

# Visitor

`Visitor` permite agregar operaciones a una estructura de objetos sin modificar los tipos que forman esa estructura.

En el ejemplo, `Circle` y `Rectangle` son figuras. `AreaVisitor` agrega la operación de calcular áreas y cada figura le indica al visitor qué tipo concreto recibió.

Este patrón puede ser útil cuando la estructura de tipos es estable, pero aparecen muchas operaciones nuevas sobre ella.

Tiene un costo importante: cada vez que se agrega un tipo nuevo de figura, también hay que actualizar `ShapeVisitor`. Por eso conviene aprenderlo, pero usarlo solo cuando esa relación sea estable.

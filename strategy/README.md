# Strategy

`Strategy` encapsula varios algoritmos intercambiables detrás de una interfaz.

En el ejemplo, `Delivery` calcula el costo usando una estrategia de envío estándar o express. `Delivery` no necesita conocer la fórmula de cada estrategia.

Para agregar un nuevo tipo de envío, se crea otra implementación de `ShippingStrategy` sin modificar `Delivery`.

Es útil cuando una decisión puede resolverse de distintas formas y se quiere elegir el algoritmo en tiempo de ejecución.

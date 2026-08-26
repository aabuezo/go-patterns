# Command

`Command` representa una acción como un objeto.

En el ejemplo, `TurnOnLightCommand` representa la acción de encender una luz. El control remoto solo conoce la interfaz `Command`, por lo que no necesita conocer los detalles de `Light`.

Esto permite guardar acciones, ejecutarlas más tarde, crear historiales o implementar deshacer. En Go, una interfaz con un método suele ser suficiente para implementar este patrón.

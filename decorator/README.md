# Decorator

`Decorator` agrega responsabilidades a un objeto envolviéndolo con otro objeto que implementa la misma interfaz.

En el ejemplo, `SimpleCoffee` representa un café. `MilkDecorator` también es un `Coffee`, pero agrega la descripción y el precio de la leche.

Se pueden encadenar varios decoradores, por ejemplo leche y luego azúcar, sin modificar `SimpleCoffee`. Es útil cuando las combinaciones de funcionalidades pueden variar en tiempo de ejecución.

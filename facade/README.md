# Facade

`Facade` ofrece una interfaz simple para usar un sistema que internamente tiene varios componentes y pasos.

En el ejemplo, crear una orden requiere reservar el producto, cobrarlo y coordinar el envío. `OrderFacade` reúne esos pasos en `PlaceOrder`.

El cliente no necesita conocer `Inventory`, `Payment` ni `Shipping`. Facade resulta útil para simplificar el uso de subsistemas complejos, sin eliminar esos componentes internos.

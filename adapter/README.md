# Adapter

`Adapter` permite que una pieza de código con una interfaz existente pueda utilizarse donde se espera otra interfaz.

En el ejemplo, `Checkout` espera un `PaymentProcessor` con un método `Pay`. `LegacyBank` tiene un método diferente, `MakeTransfer`.

`BankAdapter` traduce una llamada a `Pay` en una llamada a `MakeTransfer`. Así se puede reutilizar `LegacyBank` sin modificarlo ni modificar `Checkout`.

Es útil al integrar código antiguo, librerías externas o servicios con APIs diferentes.

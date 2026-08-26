# Proxy

`Proxy` ofrece el mismo acceso que un objeto real, pero agrega control antes de delegar la operación.

En el ejemplo, `ImageProxy` retrasa la carga de `RealImage` hasta la primera llamada a `Display`. La segunda llamada reutiliza la imagen ya cargada.

Un proxy también puede controlar permisos, registrar llamadas o agregar caché. Es útil cuando crear o acceder al objeto real es costoso.

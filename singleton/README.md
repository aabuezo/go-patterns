# Singleton

`Singleton` garantiza que exista una única instancia de un componente y ofrece un punto de acceso compartido.

En el ejemplo, `GetConfig` crea la configuración una sola vez. Las llamadas posteriores devuelven el mismo puntero.

`sync.Once` hace que la inicialización ocurra una sola vez incluso si varias partes del programa llaman a `GetConfig`.

En Go conviene usar este patrón con cuidado: muchas veces una dependencia explícita es más fácil de probar. Puede ser apropiado para una configuración inmutable o un recurso verdaderamente global.

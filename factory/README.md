# Factory

`Factory` centraliza la creación de objetos y devuelve una interfaz común, sin obligar al código cliente a conocer los tipos concretos.

En el ejemplo, `NewNotification` crea una notificación por email o SMS. El código cliente solo trabaja con la interfaz `Notification`.

Es útil cuando la creación depende de una configuración, una entrada del usuario o una regla de negocio. Si la creación es trivial, una función constructora normal suele ser suficiente.

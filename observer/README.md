# Observer

`Observer` permite notificar a varios objetos cuando ocurre un cambio.

En el ejemplo, `ProductStock` notifica a todos los suscriptores cuando llega un producto. Cada `EmailSubscriber` decide cómo reaccionar a la notificación.

El stock no conoce los detalles de cada suscriptor; solo conoce la interfaz `Observer`. Es útil para eventos, suscripciones, interfaces gráficas y actualizaciones de caché.

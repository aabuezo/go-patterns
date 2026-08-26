# Chain of Responsibility

`Chain of Responsibility` pasa una solicitud por una cadena de objetos hasta que uno puede resolverla.

En el ejemplo, el soporte de primer nivel atiende problemas normales. Las solicitudes urgentes pasan al manager.

Cada handler decide si puede hacerse cargo y, si no, deja continuar la solicitud. El patrón evita que un único componente conozca todos los casos posibles y es útil para validaciones, permisos o niveles de soporte.

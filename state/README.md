# State

`State` permite que un objeto cambie su comportamiento según su estado actual, sin acumular todos los casos en un gran `if` o `switch`.

En el ejemplo, una puerta puede estar cerrada o abierta. `ClosedDoor` y `OpenDoor` definen qué ocurre cuando se intenta abrir o cerrar en cada estado.

Cuando la puerta se abre, cambia su estado a `OpenDoor`; cuando se cierra, cambia a `ClosedDoor`.

Es útil para pedidos, conexiones, reproductores o cualquier objeto cuyo comportamiento dependa claramente de su estado.

# Composite

`Composite` permite tratar objetos individuales y grupos de objetos mediante la misma interfaz.

En el ejemplo, un `File` y un `Folder` implementan `FileSystemItem`. Un archivo devuelve su propio tamaño y una carpeta suma el tamaño de sus elementos.

El cliente solo necesita llamar a `Size`, sin distinguir si está trabajando con un archivo o una carpeta. Es útil para estructuras jerárquicas como sistemas de archivos, menús o componentes gráficos.

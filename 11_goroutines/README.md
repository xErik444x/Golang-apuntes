# Goroutines y Concurrencia!

## ¿Qué es la concurrencia?
- Básicamente la concurrencia es el proceso de poder realizar multiples tareas simultáneas.


## ¿Qué es una Goroutine?
- Una goroutine es una función o método que se ejecuta de manera concurrente con otras goroutines en el mismo espacio de direcciones. - Las goroutines son más ligeras que los hilos (threads) del sistema operativo, lo que permite manejar miles de goroutines en una sola aplicación.

## ¿Cuales son las ventajas de usar Goroutines?
- **Multiples tareas simultáneas**. - **Manejo de concurrencia de multiples procesos**. - **Manejo de concurrencia de multiples archivos**.

## Iniciando una Goroutine
Para iniciar una goroutine, simplemente coloca la palabra clave `go` antes de la llamada a la función:

```go

package main

import (
    "fmt"
    "time"
)

func decirHola() {
    fmt.Println("Hola, goroutine!")
}

func main() {
    go decirHola()
    time.Sleep(1 * time.Second)  // Espera un segundo para permitir que la goroutine termine
    fmt.Println("Hola, main!")
}
```

## ejemplo con archivos
- supongamos que queremos analizar muchos archivos en simultaneo, para aprovechar las goroutines y no esperar a que todos terminen.

```go


```
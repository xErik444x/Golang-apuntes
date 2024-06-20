---
id: 11-goroutines
title: Goroutines
sidebar_label: Goroutines
previous_page: 10-interfaces
---

- [Goroutines, Concurrencia y Canales!](#goroutines-concurrencia-y-canales)
  - [¿Qué es la concurrencia?](#qué-es-la-concurrencia)
  - [¿Qué es una Goroutine?](#qué-es-una-goroutine)
    - [¿Cuales son las ventajas de usar Goroutines?](#cuales-son-las-ventajas-de-usar-goroutines)
    - [Iniciando una Goroutine](#iniciando-una-goroutine)
  - [Canales!](#canales)
    - [Creando un canal](#creando-un-canal)
    - [Enviando y Recibiendo Datos](#enviando-y-recibiendo-datos)
    - [Ejemplo de código completo](#ejemplo-de-código-completo)

# Goroutines, Concurrencia y Canales!

## ¿Qué es la concurrencia?
- Básicamente la concurrencia es el proceso de poder realizar multiples tareas simultáneas.


## ¿Qué es una Goroutine?
- Una goroutine es una función o método que se ejecuta de manera concurrente con otras goroutines en el mismo espacio de direcciones. 
- Las goroutines son más ligeras que los hilos (threads) del sistema operativo, lo que permite manejar miles de goroutines en una sola aplicación.

### ¿Cuales son las ventajas de usar Goroutines?
- **Multiples tareas simultáneas**. 
- **Manejo de concurrencia de multiples procesos**. 
- **Manejo de concurrencia de multiples archivos**.

### Iniciando una Goroutine
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

## Canales!
- Los canales son una característica que permite la comunicación entre goroutines. 
- Un canal es un tipo que facilita la transferencia de datos entre goroutines. Podés pensar en los canales como tuberías por donde las goroutines pueden enviar y recibir valores de un tipo específico.

### Creando un canal
- Para crear un canal, se utiliza la función `make` con la palabra clave `chan` seguida del `tipo de dato` que el canal va a transportar:

```go

canal := make(chan int)
```

### Enviando y Recibiendo Datos
- Puedes enviar datos a un canal utilizando la sintaxis `<-`, así: 

```go

canal <- 42  // Envia el valor 42 al canal
```

- Para recibir datos de un canal, se utiliza la misma sintaxis pero en la dirección opuesta:

```go

valor := <-canal  // Recibe un valor desde el canal
```

### Ejemplo de código completo
- En este ejemplo, la función enviarMensaje envía un mensaje al canal, y la función main recibe ese mensaje y lo imprime.

```go

package main

import (
    "fmt"
)

func enviarMensaje(canal chan string) {
    canal <- "¡Hola desde la goroutine!"
}

func main() {
    canal := make(chan string)

    go enviarMensaje(canal)

    mensaje := <-canal  // Espera a que la goroutine envíe el mensaje
    fmt.Println(mensaje)
}
```
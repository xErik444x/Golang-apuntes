---
id: 03-5-repaso
title: Repaso rapido de defer, panic y recover
sidebar_label: Defer, panic y recover
previous_page: 03-estructuras
next_page: 04-funciones
---

- [Repaso rapido de `defer`, `panic` y `recover`](#repaso-rapido-de-defer-panic-y-recover)
  - [Función `defer`](#función-defer)
  - [Función `panic`](#función-panic)
  - [Función `recover`](#función-recover)

# Repaso rapido de `defer`, `panic` y `recover`

## Función `defer`

La instrucción `defer` pospone la ejecución de una función hasta que la función contenedora termina. Útil para cerrar archivos o limpiar recursos. Las instrucciones `defer` se ejecutan en orden inverso.

**Ejemplo:**

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 4; i++ {
        defer fmt.Println("deferred", -i)
        fmt.Println("normal", i)
    }
}
```

**Salida:**

```
normal 1
normal 2
normal 3
normal 4
deferred -4
deferred -3
deferred -2
deferred -1
```

## Función `panic`

La función `panic()` detiene el flujo normal del programa. Útil para manejar errores graves.

**Ejemplo:**

```go
package main

import "fmt"

func verificarNumero(numero int) {
    if numero < 0 {
        fmt.Println("ERROR!")
        panic("verificarNumero() recibió un número negativo")
    }
    defer fmt.Println("defer: verificarNumero(", numero, ")")
    fmt.Println("Llamada: verificarNumero(", numero, ")")
}

func main() {
    verificarNumero(5)
    fmt.Println("Esto se imprimirá después de verificarNumero(5)")

    verificarNumero(-1)
    fmt.Println("Esto no se imprimirá porque el programa entró en pánico")
}


salida:
Llamada: verificarNumero( 5 )
Diferido: verificarNumero( 5 )
Esto se imprimirá después de verificarNumero(5)
¡Pánico!
panic: verificarNumero() recibió un número negativo

goroutine 1 [running]:
main.verificarNumero(0x4bc448?)
	/tmp/sandbox2908618384/prog.go:8 +0x1ac
main.main()
	/tmp/sandbox2908618384/prog.go:18 +0x66

Program exited.
```

## Función `recover`

La función `recover()` permite recuperar el control después de una llamada a `panic()`. Debe usarse dentro de una función diferida.

**Ejemplo:**

```go
package main

import "fmt"

func verificarNumero(numero int) {
    if numero < 0 {
        fmt.Println("¡Pánico!")
        panic("verificarNumero() recibió un número negativo")
    }
    defer fmt.Println("Diferido: verificarNumero(", numero, ")")
    fmt.Println("Llamada: verificarNumero(", numero, ")")
}

func main() {
    defer func() {
        handler := recover()
        if handler != nil {
            fmt.Println("main(): recuperación", handler)
        }
    }()

    verificarNumero(5)
    fmt.Println("Esto se imprimirá después de verificarNumero(5)")

    verificarNumero(-1)
    fmt.Println("Esto no se imprimirá porque el programa entró en pánico")
}

salida:
Llamada: verificarNumero( 5 )
Diferido: verificarNumero( 5 )
Esto se imprimirá después de verificarNumero(5)
¡Pánico!
main(): recuperación verificarNumero() recibió un número negativo

Si te das cuenta, el programa no crasheó sino que se recuperó del panic.
```
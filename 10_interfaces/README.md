---
id: 10-interfaces
title: Interfaces
sidebar_label: Interfaces
previous_page: 09-json
---

- [Interfaces y Tipos!](#interfaces-y-tipos)
  - [¿Que es una interface?](#que-es-una-interface)
    - [Conceptos rápidos a tener en cuenta:](#conceptos-rápidos-a-tener-en-cuenta)
    - [Ejemplo básico de interfaz](#ejemplo-básico-de-interfaz)
  - [Usando interfaces en funciones como parámetros](#usando-interfaces-en-funciones-como-parámetros)
  - [Interfaces vacías](#interfaces-vacías)
  - [Type assertions](#type-assertions)


# Interfaces y Tipos!

## ¿Que es una interface?
- una interfaz es un tipo que define un conjunto de métodos, pero sin implementar ninguno de ellos. Cualquier tipo que implemente esos métodos se considera que implementa la interfaz.

### Conceptos rápidos a tener en cuenta:
- Interfaz:
  - Una interfaz define un conjunto de métodos sin implementar ninguno.
  - Cualquier tipo que implemente esos métodos satisface la interfaz.
- Método:
  - Un método es una función que está asociada a un tipo específico.

### Ejemplo básico de interfaz

**paso 1: Definir la interfaz**
> definimos una interfaz `Informable`, que define/contiene un método Informar que devuelve un string.


```go

type Informable interface {
    Informar() string
}
```
**paso 2:  Crear una Estructura (Tipo)**
> creamos una estructura `Persona` para implementar la interfaz `Informable`.


```go

type Persona struct {
    Nombre string
}
```

**Paso 3: Implementar el método `Informar` en la estructura `Persona`**


```go

func (p Persona) Informar() string {
    return "Hola, mi nombre es " + p.Nombre
}
```

**Paso 4: Crear un objeto de tipo `Persona` y llamar al método `Informar`**


```go

func main() {
    p := Persona{Nombre: "Juan"}
    fmt.Println(p.Informar())
}
```

## Usando interfaces en funciones como parámetros
- Puedes escribir funciones que acepten interfaces como parámetros. Esto permite que las funciones trabajen con cualquier tipo que implemente la interfaz.


```go

package main

import "fmt"

// Definición de la interfaz
type Describible interface {
    Describir() string
}

// Estructura que implementa la interfaz
type Persona struct {
    Nombre string
    Edad   int
}

// Implementación del método Describir
func (p Persona) Describir() string {
    return fmt.Sprintf("Nombre: %s, Edad: %d", p.Nombre, p.Edad)
}

// Función que acepta la interfaz Describible
func ImprimirDescripcion(d Describible) {
    fmt.Println(d.Describir())
}

func main() {
    p := Persona{"Juan", 30}
    ImprimirDescripcion(p) // Imprime "Nombre: Juan, Edad: 30"
}
```

## Interfaces vacías
- La interfaz vacía `interface{}` es una interfaz que no tiene métodos. Esto significa que todos los tipos en Go implementan la interfaz vacía.


```go

func main() {
    var x interface{}
    x = 42
    fmt.Println(x)
}
// O tambien:
package main

import "fmt"

// Función que acepta la interfaz vacía
// En este ejemplo, ImprimirValor puede aceptar cualquier tipo de valor.
func ImprimirValor(v interface{}) {
    fmt.Println(v)
}

func main() {
    ImprimirValor("Hola")
    ImprimirValor(123)
    ImprimirValor(true)
}
```

## Type assertions
- Las type assertions permiten acceder al valor concreto almacenado en una variable de tipo interfaz.
- En este ejemplo, la `type assertion` se usa en una declaración `switch` para determinar el tipo concreto del valor almacenado en `v`.


```go

package main

import "fmt"

func ImprimirValor(v interface{}) {
    switch v := v.(type) {
    case string:
        fmt.Println("Es una cadena:", v)
    case int:
        fmt.Println("Es un entero:", v)
    default:
        fmt.Println("Tipo desconocido")
    }
}

func main() {
    ImprimirValor("Hola") // Es una cadena: Hola
    ImprimirValor(123)  // Es un entero: 123
    ImprimirValor(true) // Tipo desconocido
}
```
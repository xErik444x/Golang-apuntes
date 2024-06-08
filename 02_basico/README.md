---
id: 02-variables
title: variables y declaraciones
sidebar_label: Variables y Declaraciones
previous_page: 01-introduccion
next_page: 03-estructuras
---
- [Variables y declaraciones](#variables-y-declaraciones)
    - [tipos de datos variables](#tipos-de-datos-variables)
  - [Ejercicios!](#ejercicios)

# Variables y declaraciones
> En Go, las variables se pueden declarar de varias maneras. 
* Declaración con var 
```go 
var nombre string
var edad int
```
* Declaración con inicialización
```go 
var nombre string = "Nombre generico"
var edad int = 30
```
* Declaración corta con :=
```go 
nombre := "Nombre generico2"
edad := 30
```
* Declaración de múltiples variables
```go 
var (
    nombre string
    edad   int
    altura float64
)
// Es posible declarar múltiples variables al mismo tiempo usando paréntesis.
```
* Constantes
```go 
const PI = 3.14
const Nombre = "Nombre generico3"
```
### tipos de datos variables
* Las mas comunes, bool, int, float, string:
```go
  var a bool = true     // Boolean
  var b int = 5         // Integer
  var c float32 = 3.14  // Floating point number
  var d string = "Hola!"  // String
```
* Arrays
> los arrays pueden ser de una longitud constante:
```go
  var arr [3]int = [3]int{1,2,3}
  arr2 := [4]int{1,2,3,4}
```
> o también sin predefinir el tamaño (esto lo hace el compilador cuando le asignamos el valor a la variable):
```go
  arr3 := [...]int{8, 9, 10}  // Aquí no usamos `var`
  arr2 := [...]int{1,2,3,4}
```

## Ejercicios!
- recuerda visitar la pagina [Golang playground](https://go.dev/play/) para testear y hacer el codigo.
* Ejercicio 1: Declaración y asignación de variables
  - Declara una variable llamada ciudad de tipo string sin inicializarla.
  - Declara una variable llamada poblacion de tipo int y asígnale el valor 500000.
  - Utiliza la declaración corta para declarar una variable llamada pais y asígnale el valor "Argentina".
  - Devuelve los valores de las variables creadas.
  - Si tienes problemas revisa el [ejercicio 1](./ejercicios/01.go)
  <hr>
* Ejercicio 2: Declaración de múltiples variables
  - Declara múltiples variables en una sola declaración usando  paréntesis:
    - nombre de tipo string
    - edad de tipo int
    - altura de tipo float64
    - Inicializa nombre con "Juan", edad con 25 y altura con 1.75.
    - Devuelve los valores de las variables creadas.
  - Si tienes problemas revisa el [ejercicio 2](./ejercicios/02.go)
  <hr>
* Ejercicio 3: Uso de constantes
  - Declara una constante llamada PI con el valor 3.14
  - Declara una constante llamada Nombre con el valor "John Doe"
  - Devuelve los valores de las variables creadas.
  - Si tienes problemas revisa el [ejercicio 3](./ejercicios/03.go)
  <hr>
* Ejercicio 4: Arrays
  - Declara un array de enteros arr de longitud 3 e inicialízalo con los valores 1, 2, 3.
  - Usa la declaración corta para declarar e inicializar un array de enteros arr2 de longitud 4 con los valores 4, 5, 6, 7.
  - Declara un array arr3 sin predefinir su tamaño, inicializándolo con los valores 8, 9, 10.
  - Si tienes problemas revisa el [ejercicio 4](./ejercicios/04.go)

<!-- [<< Anterior: Introducción](../01_introduccion/README.md) |
[Siguiente: estructuras de control básicas >> ](../03_estructurasDeControlBasicas/README.md) -->
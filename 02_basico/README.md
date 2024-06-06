- [Variables y declaraciones](#variables-y-declaraciones)
    - [tipos de datos variables](#tipos-de-datos-variables)

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
  var arr [...]int = [...]int{1,2,3}
  arr2 := [...]int{1,2,3,4}
```
[<< Anterior: Introducción](../01_introduccion/README.md) |
[Siguiente: estructuras de control básicas >> ](../03_estructurasDeControlBasicas/README.md)
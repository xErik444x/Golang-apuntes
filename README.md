- [Qué es esta "guía"?](#qué-es-esta-guía)
- [Qué pingo es Golang?](#qué-pingo-es-golang)
  - [Para qué se usa??](#para-qué-se-usa)
  - [Ejemplo de Código en Go:](#ejemplo-de-código-en-go)
  - [Todo muy lindo, pero como lo ejecuto y uso?](#todo-muy-lindo-pero-como-lo-ejecuto-y-uso)
- [Variables y declaraciones](#variables-y-declaraciones)
    - [tipos de datos variables](#tipos-de-datos-variables)
- [estructuras de control básicas](#estructuras-de-control-básicas)
    - [Bucle FOR](#bucle-for)
    - [Bucle Range](#bucle-range)
  
---
# Qué es esta "guía"?
- esta guía está hecha para aprender desde 0 las cosas más relevantes a la hora de aprender Golang, en especial si nunca lo viste o si sos principiante.
# Qué pingo es Golang?
* es un lenguaje de programación de código abierto desarrollado por Google.
* se caracteriza por su simplicidad, eficiencia y capacidad para manejar concurrencia de manera efectiva.
    + Go cuenta con goroutines y canales que permiten manejar múltiples tareas simultáneamente de manera eficiente.
* Go es compilado y produce código máquina que se ejecuta rápidamente.
* Go es compatible con diversas plataformas, incluyendo Windows, macOS y Linux.

## Para qué se usa??
- **Desarrollo Web**: Go se utiliza ampliamente para crear servidores web y APIs gracias a su alto rendimiento y capacidad de manejar múltiples conexiones concurrentes.
- **Herramientas y Utilidades del Sistema**: Go se usa para desarrollar herramientas de línea de comandos y otros tipos de software del sistema.

## Ejemplo de Código en Go:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hola, Mundo!")
}
```

## Todo muy lindo, pero como lo ejecuto y uso?
> [basandonos en la guia oficial](https://go.dev/doc/tutorial/getting-started) luego de instalar todos los requisitos
* creamos la carpeta del proyecto
* entramos dentro de la carpeta e inicializamos el proyecto
```bash
$ go mod init example/hello
```
* en este caso estamos creando el modulo example/hello (los nombres de módulos pueden llevar varias palabras separadas con `/`)
* una vez creado el modulo, debemos crear el archivo que contenga la función `main` (principal) del programa.
* dentro de un archivo con terminación .go ej: principal.go escribimos:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hola Manu mirá estoy hackeando la nasa!!!")
}
```
* una vez guardado solo resta ejecutarlo con `go run .` o `go run principal.go`
* y listo, ahí tendrías tu primer programa en Golang!

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
> o también sin predefinir el tamaño:
```go
  var arr [...]int = [...]int{1,2,3}
  arr2 := [...]int{1,2,3,4}
```
# estructuras de control básicas
### Bucle FOR
* el for de Go no es muy diferente si miramos al de Java o Javascript 
* Tenemos 3 parametros
  * el valor inicial
  * la condición
  * operación al valor inicial
```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
    fmt.Println("valor de i: ", i)
    // valor de i:  0
    // valor de i:  1
    // ....
}
fmt.Println("total: ", sum)
// total:  45
```

### Bucle Range
* El `range` es utilizado para iterar `arrays`, `slices` y `maps` (tambien `canales`, pero todavía es pronto) en Go utilizando for:
* Tenemos 3 parámetros
  * 
  * el array
  * otro valor
```go
numbers := []int{1, 2, 3, 4, 5}
sum := 0
for _, number := range numbers {
    sum += number
}
fmt.Println("La suma de los números es:", sum)
```
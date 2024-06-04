- [que pingo es Golang?](#que-pingo-es-golang)
  - [para qué se usa??](#para-qué-se-usa)
  - [Ejemplo de Código en Go:](#ejemplo-de-código-en-go)
  - [Todo muy lindo, pero como lo ejecuto y uso?](#todo-muy-lindo-pero-como-lo-ejecuto-y-uso)
- [variables y declaraciones](#variables-y-declaraciones)
  
---

# que pingo es Golang?
* es un lenguaje de programación de código abierto desarrollado por Google.
* se caracteriza por su simplicidad, eficiencia y capacidad para manejar concurrencia de manera efectiva.
    + Go cuenta con goroutines y canales que permiten manejar múltiples tareas simultáneamente de manera eficiente.
* Go es compilado y produce código máquina que se ejecuta rápidamente.
* Go es compatible con diversas plataformas, incluyendo Windows, macOS y Linux.

## para qué se usa??
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

# variables y declaraciones
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
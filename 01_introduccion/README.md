---
id: 01-introduccion
title: Introducción
sidebar_label: Introducción
next_page: 02-variables
---
# Índice
- [Índice](#índice)
- [Qué pingo es Golang?](#qué-pingo-es-golang)
  - [Para qué sirve??](#para-qué-sirve)
  - [Ejemplo de Código en Go](#ejemplo-de-código-en-go)
  - [Todo muy lindo, pero como lo ejecuto y uso?](#todo-muy-lindo-pero-como-lo-ejecuto-y-uso)
  - [Ejercicios!](#ejercicios)
# Qué pingo es Golang?
* es un lenguaje de programación de código abierto desarrollado por Google.
* se caracteriza por su simplicidad, eficiencia y capacidad para manejar concurrencia de manera efectiva.
    + Go cuenta con goroutines y canales que permiten manejar múltiples tareas simultáneamente de manera eficiente.
* Go es compilado y produce código máquina que se ejecuta rápidamente.
* Go es compatible con diversas plataformas, incluyendo Windows, macOS y Linux.

## Para qué sirve??
- **Desarrollo Web**: Go se utiliza ampliamente para crear servidores web y APIs gracias a su alto rendimiento y capacidad de manejar múltiples conexiones concurrentes.
- **Herramientas y Utilidades del Sistema**: Go se usa para desarrollar herramientas de línea de comandos y otros tipos de software del sistema.

## Ejemplo de Código en Go

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
* en este caso estamos creando el modulo example/hello (los nombres de módulos pueden llevar varias palabras separadas con `/` lo ideal es prefix/texto_descriptivo)
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

## Ejercicios!
* Crear un programa que imprima "Hola, Mundo!"
  * pasos:
  * crear carpeta
  * inicializar el proyecto
  * crear el archivo principal
  * escribir el programa
  * ejecutar el programa!


[Siguiente: Conceptos Básicos / variables >> ](../02_basico/README.md)
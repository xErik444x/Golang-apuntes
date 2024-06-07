- [Introducción a la Creación de CLI y Compilaciones](#introducción-a-la-creación-de-cli-y-compilaciones)
  - [Uso de la Biblioteca flag](#uso-de-la-biblioteca-flag)
    - [Quiero generar un ejecutable del programa, cómo hago?](#quiero-generar-un-ejecutable-del-programa-cómo-hago)
    - [compilación cruzada:](#compilación-cruzada)
      - [Listar las plataformas disponibles:](#listar-las-plataformas-disponibles)
      - [Compilar:](#compilar)
      - [compilar en Linux:](#compilar-en-linux)
      - [compilar en Windows:](#compilar-en-windows)
      - [compilar en Macos:](#compilar-en-macos)
  - [Ejercicios!](#ejercicios)
    - [Crear un CLI que multiplique dos números](#crear-un-cli-que-multiplique-dos-números)

# Introducción a la Creación de CLI y Compilaciones
- Los CLI (Command Line Interfaces) son una herramienta para interactuar con aplicaciones a través de la línea de comandos. En Go, crear CLI es eficiente y estructurado. 
- Nosotros Vamos a ver cómo crear un CLI básico utilizando la biblioteca flag, que es parte de la biblioteca estándar de Go.
  
## Uso de la Biblioteca flag
- La biblioteca flag de Go permite definir opciones y argumentos de línea de comandos para los programas.
ejemplo de código con la biblioteca flag:
```go
package main

import (
    "flag"
    "fmt"
)

func main() {
    // Definimos un parámetro (flag) de tipo string
    // tiene 3 parámetros:
    // - nombre del parámetro
    // - valor por defecto
    // - descripción del parámetro
    nombre := flag.String("nombre", "Mundo", "un nombre para saludar")
    edad := flag.Int("edad", 1, "edad")

    // Parseamos los flags / parámetros
    flag.Parse()

    // Usamos el valor del flag con * porque es un puntero, no podemos modificarlo directamente.
    fmt.Printf("Hola, %s!\n", *nombre)
    fmt.Printf("Edad: %d\n", *edad)
}
```
- Cómo lo ejecuto con parámetros?
- `go run main.go --nombre=Manuel`
- si no colocamos un valor por defecto, el valor por defecto es `Mundo`

### Quiero generar un ejecutable del programa, cómo hago?
- para crear un ejecutable, necesitamos la biblioteca `go build`
- con colocar `go build -o saludo` generamos un .exe (en windows, si usas otra plataforma el tipo de archivo puede ser diferente) `saludo`
- para ejecutarlo con `./saludo --nombre=Juan`, o `saludo` y sus parámetros

### compilación cruzada:
- quiero generar un ejecutable del programa, para otras plataformas (macos, windows, etc) cómo hago?

#### Listar las plataformas disponibles:
- `go tool dist list`
  
#### Compilar:
- Para construir el ejecutable para una plataforma diferente, puedes usar las variables de entorno `GOOS` y `GOARCH` para especificar el sistema operativo y la arquitectura

#### compilar en Linux:
- `GOOS=linux GOARCH=amd64 go build -o saludo-linux-amd64 saludo.go`

#### compilar en Windows:
- `GOOS=windows GOARCH=amd64 go build -o saludo-windows-amd64.exe saludo.go`
  
#### compilar en Macos:
- `GOOS=darwin GOARCH=arm64 go build -o saludo-darwin-arm64 saludo.go`

## Ejercicios!
### Crear un CLI que multiplique dos números
- crear un CLI en Go que tome dos números como parámetros y devuelva su producto. 
- Si no sabes como hacerlo, revisa el [ejercicio 1 main.go](ejercicios/01/main.go)

[<< Anterior: Type Casting \& Maps \& Slices \& Structs \& Pointers](../05_typeCastingMapsSlicesStructsPointers/README.md)
